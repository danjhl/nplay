package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCreatesPlaylist(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "tmp")
	defer os.Remove(tmpDir)

	playlist := tmpDir + "/playlist.txt"
	createFiles(tmpDir, "dummy1.mp3", "dummy2.mp3", "dummy3.mp3")

	var buffer bytes.Buffer
	err := Run(&buffer, "add", []string{tmpDir})
	assert.Nil(t, err)

	data, err := os.ReadFile(playlist)
	assert.Nil(t, err)

	expected := "dummy1.mp3\n" +
		"dummy2.mp3\n" +
		"dummy3.mp3\n"

	assert.Equal(t, expected, string(data))
}

func TestAddAddsFilesToExistingPlaylist(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "tmp")
	defer os.Remove(tmpDir)

	playlist := tmpDir + "/playlist.txt"
	createFiles(tmpDir, "dummy1.mp3", "dummy2.mp3", "dummy3.mp3")
	createPlaylist(tmpDir, "dummy1.mp3")

	var buffer bytes.Buffer
	err := Run(&buffer, "add", []string{tmpDir})
	assert.Nil(t, err)

	data, err := os.ReadFile(playlist)
	assert.Nil(t, err)

	expected := "dummy1.mp3\n" +
		"dummy2.mp3\n" +
		"dummy3.mp3\n"

	assert.Equal(t, expected, string(data))
}

func TestAddOnlyAddsSupportedFiles(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "tmp")
	defer os.Remove(tmpDir)

	playlist := tmpDir + "/playlist.txt"
	createFiles(tmpDir, "dummy1.ogg", "dummy2.wav", "dummy3.mp3", "dummy.txt", "dummy.mp4")

	var buffer bytes.Buffer
	err := Run(&buffer, "add", []string{tmpDir})
	assert.Nil(t, err)

	data, err := os.ReadFile(playlist)
	assert.Nil(t, err)

	expected := "dummy1.ogg\n" +
		"dummy2.wav\n" +
		"dummy3.mp3\n"

	assert.Equal(t, expected, string(data))
}

func createFiles(dir string, files ...string) {
	for _, file := range files {
		f, err := os.Create(dir + "/" + file)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

func createPlaylist(dir string, entries ...string) {
	f := dir + "/playlist.txt"
	playlist, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	defer playlist.Close()

	for _, entry := range entries {
		playlist.WriteString(entry + "\n")
	}
}
