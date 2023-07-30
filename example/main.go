package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	h, w := termbox.Size()
	termbox.Close()
	fmt.Println(h, w)
}
