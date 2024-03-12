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

func (h Help) Execute(args []string, out io.Writer) error {
	panic("unimplemented")
}

func (h Help) ExecuteHelp(args []string, out io.Writer, cmds []Cmd) {
	builder := strings.Builder{}

	for i, cmd := range cmds {
		help := cmd.Help()
		if help == "" {
			continue
		}
		builder.WriteString(help)
		if i < len(cmds)-1 {
			builder.WriteString("\n")
		}
	}
	out.Write([]byte(builder.String()))
}

func (h Help) Help() string {
	return ""
}
