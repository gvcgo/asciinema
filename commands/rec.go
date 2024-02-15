package commands

import (
	"github.com/gvcgo/asciinema/asciicast"
)

type RecordCommand struct {
	Env      map[string]string
	Recorder asciicast.Recorder
}

func NewRecordCommand(env map[string]string) *RecordCommand {
	return &RecordCommand{
		Env:      env,
		Recorder: asciicast.NewRecorder(),
	}
}

func (c *RecordCommand) Execute(command, title string, assumeYes bool, maxWait float64) (asciicast.Asciicast, error) {
	return c.Recorder.Record(command, title, maxWait, assumeYes, c.Env)
}
