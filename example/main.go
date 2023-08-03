package main

import (
	"os"

	"github.com/moqsien/asciinema/cmd"
)

func main() {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	cli := cmd.New(args...)
	cli.Play()
}
