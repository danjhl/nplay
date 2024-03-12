package main

import "io"

type Cmd interface {
	Name() string
	Help() string
	Execute(args []string, out io.Writer) error
}
