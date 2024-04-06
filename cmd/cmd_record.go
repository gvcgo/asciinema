package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"runtime"
	"strings"

	"github.com/gvcgo/asciinema/asciicast"
	"github.com/gvcgo/asciinema/commands"
	"github.com/gvcgo/asciinema/util"
	"github.com/olivere/ndjson"
)

var descardingList []string = []string{
	`?\u001b\\\u001b[6n`,
}

func verify(line string) bool {
	for _, s := range descardingList {
		if strings.Contains(line, s) {
			return false
		}
	}
	return true
}

func FixCast(fPath string) {
	content, _ := os.ReadFile(fPath)
	if len(content) > 0 {
		sList := strings.Split(string(content), "\n")
		data := []string{}
		for _, line := range sList {
			if verify(line) {
				data = append(data, line)
			}
		}
		if len(data) > 0 {
			s := strings.Join(data, "\n")
			os.WriteFile(fPath, []byte(s), os.ModePerm)
		}
	}
}

func (r *Runner) Rec() error {
	command := "C:\\WINDOWS\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	if ok, _ := util.PathIsExist(command); !ok {
		command = "powershell.exe"
	}
	if runtime.GOOS != "windows" {
		command = util.FirstNonBlank(os.Getenv("SHELL"), cfg.RecordCommand())
	}

	if r.Quite {
		util.BeQuiet()
		r.AssumeYes = true
	}

	cmd := commands.NewRecordCommand(env)
	cast, err := cmd.Execute(command, r.Title, r.AssumeYes, r.MaxWait)
	if err != nil {
		return err
	}
	r.Cast = &cast
	var buf bytes.Buffer
	result := ndjson.NewWriter(&buf)

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
		return err
	}
	for _, f := range cast.Stdout {
		if err := result.Encode([]interface{}{f.Time, "o", string(f.EventData)}); err != nil {
			panic(err)
		}
	}

	err = os.WriteFile(r.FilePath, buf.Bytes(), os.ModePerm)
	if err == nil {
		FixCast(r.FilePath)
	}
	return err
}
