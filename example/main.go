package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/moqsien/asciinema"
// 	"github.com/moqsien/asciinema/asciicast"
// )

// var Header string = `{"version":2,"width":%d,"height":%d,"timestamp":%d,"env":{"TERM":"%s","SHELL":"%s"}}`

// func writeFile(ac *asciicast.Asciicast) {
// 	content := []string{}
// 	content = append(content, fmt.Sprintf(Header, ac.Width, ac.Height, ac.Timestamp, ac.Env.Term, ac.Env.Shell))
// 	for _, v := range ac.Stdout {
// 		if b, err := v.MarshalJSON(); err == nil {
// 			content = append(content, string(b))
// 		}
// 	}
// 	os.WriteFile("./play_test.txt", []byte(strings.Join(content, "\n")), os.ModePerm)
// }

// func main() {
// 	cli := asciinema.New()
// 	cast, _, _ := cli.Rec()
// 	writeFile(cast)
// 	fmt.Println(cast.Stdout)
// 	fmt.Println(cast.Timestamp)
// 	fmt.Println(cast.Width)
// 	fmt.Println(cast.Height)
// 	fmt.Println(cast.Env.Shell)
// 	fmt.Println(cast.Env.Term)
// }

import (
	"context"
	"io"
	"log"
	"os"

	conpty "github.com/EgeBalci/conpty-go"
	"github.com/nsf/termbox-go"
)

func main() {
	commandLine := `c:\windows\system32\cmd.exe`
	termbox.Init()
	w, h := termbox.Size()
	termbox.Close()
	opt := conpty.ConPtyDimensions(w, h)
	cpty, err := conpty.Start(commandLine, opt)
	if err != nil {
		log.Fatalf("Failed to spawn a pty:  %v", err)
	}
	defer cpty.Close()

	go func() {
		go io.Copy(os.Stdout, cpty)
		io.Copy(cpty, os.Stdin)
	}()

	exitCode, err := cpty.Wait(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("ExitCode: %d", exitCode)
}
