//go:build darwin || freebsd || dragonfly || linux

package asciicast

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gvcgo/asciinema/util"
)

func (r *AsciicastRecorder) checkTerminalSize() chan<- bool {
	rows, cols, _ := r.Terminal.Size()
	doneChan := make(chan bool)
	go func() {
		winch := make(chan os.Signal, 1)
		signal.Notify(winch, syscall.SIGWINCH)

		defer signal.Stop(winch)
		defer close(winch)
		defer close(doneChan)

		for {
			select {
			case <-winch:
				newRows, newCols, _ := r.Terminal.Size()
				if cols != newCols || rows != newRows {
					cols, rows = newCols, newRows
					util.ReplaceWarningf("Current terminal size is %s.", fmt.Sprintf("%dx%d", cols, rows))
				}
			case <-doneChan:
				return
			}
		}
	}()
	return doneChan
}
