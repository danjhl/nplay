package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelp(t *testing.T) {
	var out bytes.Buffer
	Run(&out, "help", []string{})

	expected := Add{}.Help() + "\n\n" + Update{}.Help() + "\n\n"
	assert.Equal(t, expected, out.String())
}
