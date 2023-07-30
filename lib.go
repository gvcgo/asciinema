package asciinema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/moqsien/asciinema/asciicast"
	"github.com/moqsien/asciinema/commands"
	"github.com/moqsien/asciinema/util"
	"github.com/olivere/ndjson"
)

// Options options to pass to various commands.
// These are common flags passed to the asciinema cli.
type Options struct {
	Title   string
	MaxWait float64
	Yes     bool
	Quite   bool
}

// New creates a new Options instance.
func New(opts ...Options) *Options {
	var o Options
	if len(opts) == 0 {
		return &Options{
			Title:   "",
			MaxWait: 1.0,
			Yes:     false,
			Quite:   false,
		}
	}

	options := opts[0]

	o = options
	return &o
}

// Play plays the given asciicast. Use asciicast.Asciicast to unmarshal
// read from the asciicast file.
func (o *Options) Play(cast *asciicast.Asciicast) error {
	initAsciinema()
	cmd := commands.NewPlayCommand()
	return cmd.Execute(cast, o.MaxWait)
}

// Rec records the terminal and returns the asciicast and error.
func (o *Options) Rec() (*asciicast.Asciicast, *bytes.Buffer, error) {
	initAsciinema()
	command := util.FirstNonBlank(os.Getenv("SHELL"), cfg.RecordCommand())
	title := o.Title
	assumeYes := o.Yes

	if o.Quite {
		util.BeQuiet()
		assumeYes = true
	}

	maxWait := o.MaxWait
	cmd := commands.NewRecordCommand(env)
	cast, err := cmd.Execute(command, title, assumeYes, maxWait)
	if err != nil {
		return &asciicast.Asciicast{}, nil, err
	}

	var buf bytes.Buffer
	r := ndjson.NewWriter(&buf)

	header := &asciicast.Header{
		Version:   cast.Version,
		Command:   cast.Command,
		Title:     cast.Title,
		Width:     cast.Width,
		Height:    cast.Height,
		Timestamp: cast.Timestamp,
		Duration:  cast.Duration,
		Env:       cast.Env,
	}

	// add header
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&header); err != nil {
		return &asciicast.Asciicast{}, nil, err
	}
	for _, f := range cast.Stdout {
		if err := r.Encode([]interface{}{f.Time, "o", string(f.EventData)}); err != nil {
			panic(err)
		}
	}

	return &cast, &buf, nil
}

func initAsciinema() {
	env = environment()

	if !util.IsUtf8Locale(env) {
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

	cfg, err = util.GetConfig(env)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const Version = "1.2.0"

func environment() map[string]string {
	env := map[string]string{}

	for _, keyval := range os.Environ() {
		pair := strings.SplitN(keyval, "=", 2)
		env[pair[0]] = pair[1]
	}

	return env
}

func showCursorBack() {
	fmt.Fprintf(os.Stdout, "\x1b[?25h")
}

var (
	env map[string]string
	cfg *util.Config
	err error
)
