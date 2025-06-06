package tools

import (
	"io"
	"os"

	"github.com/fatih/color"
)

type ColorWriter struct {
	w io.Writer
	c *color.Color
}

func NewErrWriter() *ColorWriter {
	return &ColorWriter{
		w: os.Stderr,
		c: color.New(color.FgRed),
	}
}

func (cw *ColorWriter) Write(p []byte) (n int, err error) {
	cw.c.SetWriter(cw.w)
	defer cw.c.UnsetWriter(cw.w)
	return cw.c.Print(string(p))
}
