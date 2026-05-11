package console

import (
	"fmt"
	"io"
)

type IO interface {
	Println(text string)
}

type StdIO struct {
	writer io.Writer
}

func NewStdIO(writer io.Writer) *StdIO {
	return &StdIO{writer: writer}
}

func (s *StdIO) Println(text string) {
	_, _ = fmt.Fprintln(s.writer, text)
}
