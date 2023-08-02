package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/moqsien/asciinema"
	"github.com/moqsien/asciinema/asciicast"
)

var Header string = `{"version":2,"width":%d,"height":%d,"timestamp":%d,"env":{"TERM":"%s","SHELL":"%s"}}`

func writeFile(ac *asciicast.Asciicast) {
	content := []string{}
	content = append(content, fmt.Sprintf(Header, ac.Width, ac.Height, ac.Timestamp, ac.Env.Term, ac.Env.Shell))
	for _, v := range ac.Stdout {
		if b, err := v.MarshalJSON(); err == nil {
			content = append(content, string(b))
		}
	}
	os.WriteFile("./play_test.txt", []byte(strings.Join(content, "\n")), os.ModePerm)
}

func main() {
	cli := asciinema.New()
	cast, _, _ := cli.Rec()
	writeFile(cast)
}
