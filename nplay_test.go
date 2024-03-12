package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelp(t *testing.T) {
	var out bytes.Buffer
	err := Run(&out, "help", []string{})

	expected := "test"
	assert.Nil(t, err)
	assert.Equal(t, expected, out.String())
}

func TestInit(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}
