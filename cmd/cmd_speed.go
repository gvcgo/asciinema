package cmd

import (
	"github.com/gvcgo/asciinema-edit/cast"
	"github.com/gvcgo/asciinema-edit/commands/transformer"
	"github.com/gvcgo/asciinema-edit/editor"
)

type speedTransformation struct {
	from   float64
	to     float64
	factor float64
}

func (t *speedTransformation) Transform(c *cast.Cast) (err error) {
	if t.from == 0 && t.to == 0 {
		t.from = c.EventStream[0].Time
		t.to = c.EventStream[len(c.EventStream)-1].Time
	}

	err = editor.Speed(c, t.factor, t.from, t.to)
	return
}

func (r *Runner) Speed(inFilePath, outFilePath string, factor, start, end float64) error {
	transformation := &speedTransformation{
		factor: factor,
		from:   start,
		to:     end,
	}
	t, err := transformer.New(transformation, inFilePath, outFilePath)
	if err != nil {
		return err
	}
	defer t.Close()
	return t.Transform()
}
