//go:build darwin || freebsd || dragonfly || linux

package terminal

import (
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	"unsafe"

	"github.com/creack/pty"
	"github.com/creack/termios/raw"
	"github.com/gvcgo/asciinema/util"

	// "golang.org/x/crypto/ssh/terminal"
	terminal "golang.org/x/term"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type Pty struct {
	Stdin  *os.File
	Stdout *os.File
}

func NewTerminal() Terminal {
	return &Pty{Stdin: os.Stdin, Stdout: os.Stdout}
}

func (p *Pty) Size() (int, int, error) {
	return pty.Getsize(p.Stdout)
}

func (p *Pty) Record(command string, w io.Writer, envs ...string) error {
	// start command in pty
	cmd := exec.Command("sh", "-c", command)

	if len(envs) == 0 {
		envs = append(os.Environ(), "ASCIINEMA_REC=1")
	}

	cmd.Env = envs
	master, err := pty.Start(cmd)
	if err != nil {
		return err
	}
	defer master.Close()

	// install WINCH signal handler
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGWINCH)
	defer signal.Stop(signals)
	go func() {
		for range signals {
			p.resize(master)
		}
	}()
	defer close(signals)

	// put stdin in raw mode (if it's a tty)
	fd := p.Stdin.Fd()
	if terminal.IsTerminal(int(fd)) {
		oldState, err := raw.MakeRaw(fd)
		if err != nil {
			return err
		}
		defer raw.TcSetAttr(fd, oldState)
	}

	// do initial resize
	p.resize(master)

	// start stdin -> master copying
	stop := util.Copy(master, p.Stdin)

	// copy pty master -> p.stdout & w

	stdout := transform.NewWriter(w, unicode.UTF8.NewEncoder())
	defer stdout.Close()

	stdoutWaitChan := make(chan struct{})
	go func() {
		io.Copy(io.MultiWriter(p.Stdout, stdout), master)
		stdoutWaitChan <- struct{}{}
	}()

	// wait for the process to exit and reap it
	cmd.Wait()

	// wait for master -> stdout copying to finish
	// sometimes after process exits reading from master blocks forever (race condition?)
	// we're using timeout here to overcome this problem
	select {
	case <-stdoutWaitChan:
	case <-time.After(200 * time.Millisecond):
	}

	// stop stdin -> master copying
	stop()

	return nil
}

func (p *Pty) Write(data []byte) error {
	_, err := p.Stdout.Write(data)
	if err != nil {
		return err
	}

	err = p.Stdout.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (p *Pty) resize(f *os.File) {
	var rows, cols int

	if terminal.IsTerminal(int(p.Stdout.Fd())) {
		rows, cols, _ = p.Size()
	} else {
		rows = 24
		cols = 80
	}

	Setsize(f, rows, cols)
}

type winsize struct {
	ws_row    uint16
	ws_col    uint16
	ws_xpixel uint16
	ws_ypixel uint16
}

func ioctl(fd, cmd, ptr uintptr) error {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr)
	if e != 0 {
		return e
	}
	return nil
}

func Setsize(f *os.File, rows int, cols int) error {
	var ws winsize
	ws.ws_row = uint16(rows)
	ws.ws_col = uint16(cols)
	return ioctl(f.Fd(), syscall.TIOCSWINSZ, uintptr(unsafe.Pointer(&ws)))
}
