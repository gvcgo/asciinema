//go:build windows

package terminal

import (
	"context"
	"io"
	"os"

	"github.com/moqsien/asciinema/util"
	"github.com/nsf/termbox-go"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//TODO: https://github.com/marcomorain/go-conpty

type Pty struct {
	Stdin  *os.File
	Stdout *os.File
}

func NewTerminal() Terminal {
	return &Pty{Stdin: os.Stdin, Stdout: os.Stdout}
}

func (p *Pty) Size() (int, int, error) {
	if err := termbox.Init(); err != nil {
		return 0, 0, err
	}
	defer termbox.Close()
	w, h := termbox.Size()
	return w, h, nil
}

// command = "powershell.exe"
func (p *Pty) Record(command string, w io.Writer) error {
	width, height, _ := p.Size()
	if width == 0 {
		width = 180
	}
	if height == 0 {
		height = 100
	}
	opt := util.ConPtyDimensions(width, height)
	cpty, err := util.Start(command, opt)
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

	_, err = cpty.Wait(context.Background())
	if err != nil {
		return err
	}
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
