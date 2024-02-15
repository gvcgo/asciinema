package main

import (
	"os"

	"github.com/gvcgo/asciinema/cmd"
)

func main() {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	cli := cmd.New(args...)
	// fmt.Println(cli.Auth())
	// cli.Rec()
	cli.Play()
	// info, _ := cli.Upload()
	// fmt.Println(info)
}
