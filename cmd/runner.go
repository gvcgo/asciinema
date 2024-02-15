package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gvcgo/asciinema/asciicast"
	"github.com/gvcgo/asciinema/util"
)

type Runner struct {
	Title     string
	MaxWait   float64
	AssumeYes bool
	Quite     bool
	FilePath  string
	Cast      *asciicast.Asciicast
}

func New(filename ...string) (r *Runner) {
	r = &Runner{
		Title:     "asciinema_default",
		MaxWait:   1.0,
		AssumeYes: false,
		Quite:     false,
		FilePath:  "asciinema_default.cast",
	}
	if len(filename) > 0 {
		r.FilePath = filename[0]
		name := filepath.Base(r.FilePath)
		r.Title = strings.Split(name, ".")[0]
	}
	initAsciinema()
	return
}

/*
Envs
*/
const (
	Version = "1.2.0"
)

var (
	cfg *util.Config
	env map[string]string
)

func showCursorBack() {
	fmt.Fprintf(os.Stdout, "\x1b[?25h")
}

func initAsciinema() {
	env = map[string]string{}
	for _, keyval := range os.Environ() {
		pair := strings.SplitN(keyval, "=", 2)
		env[pair[0]] = pair[1]
	}

	if runtime.GOOS != "windows" && !util.IsUtf8Locale(env) {
		fmt.Println("asciinema needs a UTF-8 native locale to run. Check the output of `locale` command.")
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		showCursorBack()
		os.Exit(1)
	}()
	defer showCursorBack()

	var err error
	cfg, err = util.GetConfig(env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
