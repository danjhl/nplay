package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelp(t *testing.T) {
	var out bytes.Buffer
	Run(&out, "help", []string{})

	expected := ""
	assert.Equal(t, expected, out.String())
}
