//go:build windows

package asciicast

import (
	"fmt"
	"time"

	"github.com/gvcgo/asciinema/util"
)

func (r *AsciicastRecorder) checkTerminalSize() chan<- bool {
	rows, cols, _ := r.Terminal.Size()
	doneChan := make(chan bool)
	go func() {
		defer close(doneChan)
		for {
			select {
			case <-doneChan:
				return
			default:
				newRows, newCols, _ := r.Terminal.Size()
				if cols != newCols || rows != newRows {
					cols, rows = newCols, newRows
					util.ReplaceWarningf("Current terminal size is %s.", fmt.Sprintf("%dx%d", cols, rows))
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	return doneChan
}
