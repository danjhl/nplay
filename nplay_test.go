package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownCommand(t *testing.T) {
	var out bytes.Buffer
	err := Run(&out, "unknown", []string{})
	assert.NotNil(t, err)
	assert.Equal(t, "Unknown command: 'unknown'", err.Error())
}
