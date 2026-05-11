package testsupport

import "strings"

type TestConsole struct {
	output strings.Builder
}

func NewTestConsole(_ ...string) *TestConsole {
	return &TestConsole{}
}

func (t *TestConsole) Println(text string) {
	t.output.WriteString(text)
	t.output.WriteString("\n")
}

func (t *TestConsole) Complete(line string) []string {
	return nil
}

func (t *TestConsole) Output() string {
	return t.output.String()
}
