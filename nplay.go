package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	out := os.Stdout
	args := []string{}
	cmdName := "help"

	if len(os.Args) >= 2 {
		args = os.Args[1:]
		cmdName = os.Args[1]
	}
	Run(out, cmdName, args)
}

func Run(out io.Writer, cmdName string, args []string) error {
	var help = Help{}
	var cmds []Cmd = []Cmd{help}

	for _, cmd := range cmds {
		if cmd.Name() == cmdName {
			if cmd == help {
				(cmd.(Help)).ExecuteHelp(args, out, cmds)
				return nil
			}
			return cmd.Execute(args, out)
		}
	}
	return errors.New("Unknown command: '" + cmdName + "'")
}

func handle(err error, out io.Writer) {
	if err != nil {
		fmt.Fprintf(out, "Error: %s\n", err.Error())
	}
}
