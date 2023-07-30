package commands

import "github.com/moqsien/asciinema/asciicast"

type PlayCommand struct {
	Player asciicast.Player
}

func NewPlayCommand() *PlayCommand {
	return &PlayCommand{
		Player: asciicast.NewPlayer(),
	}
}

func (c *PlayCommand) Execute(cast *asciicast.Asciicast, maxWait float64) error {
	return c.Player.Play(cast, maxWait)
}
