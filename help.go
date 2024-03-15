package main

import (
	"io"
	"strings"
)

type Help struct {
}

func (h Help) Name() string {
	return "help"
}

func (h Help) Execute(args []string) error {
	panic("unimplemented")
}

func (h Help) ExecuteHelp(args []string, out io.Writer, cmds []Cmd) {
	builder := strings.Builder{}

	for _, cmd := range cmds {
		help := cmd.Help()
		if help == "" {
			continue
		}
		builder.WriteString(help)
		builder.WriteString("\n\n")
	}
	out.Write([]byte(builder.String()))
}

func (h Help) Help() string {
	return ""
}
