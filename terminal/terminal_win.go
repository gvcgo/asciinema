//go:build windows

package terminal

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/gvcgo/asciinema/util/winpty"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

/*
https://github.com/marcomorain/go-conpty
https://github.com/ActiveState/termtest
*/

type Pty struct {
	Stdin  *os.File
	Stdout *os.File
}

func NewTerminal() Terminal {
	return &Pty{Stdin: os.Stdin, Stdout: os.Stdout}
}

func (p *Pty) Size() (rows, cols int, err error) {
	coord, err := winpty.WinConsoleScreenSize()
	return coord.Y, coord.X, err
}

func (p *Pty) Record(command string, w io.Writer) error {
	height, width, _ := p.Size()
	if width == 0 {
		width = 180
	}
	if height == 0 {
		height = 100
	}

	envs := append(os.Environ(), "ASCIINEMA_REC=1")
	if len(TerminalEnvs) > 0 {
		envs = TerminalEnvs
	}

	// winpty.EnableVirtualTerminalProcessing()
	cpty, err := winpty.Start(command, &winpty.COORD{X: width, Y: height}, envs)
	if err != nil {
		return err
	}
	defer cpty.Close()

	stdout := transform.NewWriter(w, unicode.UTF8.NewEncoder())
	defer stdout.Close()

	go func() {
		go io.Copy(io.MultiWriter(p.Stdout, stdout), cpty)
		io.Copy(cpty, p.Stdin)
	}()

	exitCode, err := cpty.Wait(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("ExitCode: %d", exitCode)
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
