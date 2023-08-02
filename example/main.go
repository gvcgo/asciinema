package main

import (
	"os"

	"github.com/moqsien/asciinema/commands"
)

func main() {
	cli := commands.New()
	_, b, _ := cli.Rec()
	os.WriteFile("./test_test.txt", b.Bytes(), os.ModePerm)
}
