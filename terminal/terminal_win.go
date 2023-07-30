//go:build windows

package terminal

import (
	"io"
	"os"

	"github.com/nsf/termbox-go"
)

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

func (p *Pty) Record(command string, w io.Writer) error {
	return nil
}

func (p *Pty) Write(data []byte) error {
	return nil
}
