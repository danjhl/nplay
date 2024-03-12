package main

import "io"

func main() {

}

func Run(out io.Writer, cmd string, args []string) error {
	if cmd == "help" {
		help(out)
	}
	return nil
}

func help(out io.Writer) {
	out.Write([]byte("test"))
}
