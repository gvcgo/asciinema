package cmd

import (
	"github.com/gvcgo/asciinema-edit/cast"
	"github.com/gvcgo/asciinema-edit/commands/transformer"
	"github.com/gvcgo/asciinema-edit/editor"
)

// Cut: Removes a certain range of time frames.
type cutTransformation struct {
	from float64
	to   float64
}

func (t *cutTransformation) Transform(c *cast.Cast) (err error) {
	err = editor.Cut(c, t.from, t.to)
	return
}

func (r *Runner) Cut(inFilePath, outFilePath string, start, end float64) error {
	transformation := &cutTransformation{
		from: start,
		to:   end,
	}
	t, err := transformer.New(transformation, inFilePath, outFilePath)
	if err != nil {
		return err
	}
	defer t.Close()
	err = t.Transform()
	if err == nil {
		FixHeaderForEditOperations(inFilePath, outFilePath)
	}
	return err
}
