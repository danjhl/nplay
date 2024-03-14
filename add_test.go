package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCreatesPlaylist(t *testing.T) {
	playlist := "./test/playlist.txt"
	os.Mkdir("test", os.ModePerm)
	createFiles("test", "dummy1.mp3", "dummy2.mp3", "dummy3.mp3")

	var buffer bytes.Buffer
	err := Run(&buffer, "add", []string{"test"})
	assert.Nil(t, err)

	data, err := os.ReadFile(playlist)
	assert.Nil(t, err)

	expected := "dummy1.mp3\n" +
		"dummy2.mp3\n" +
		"dummy3.mp3\n"

	assert.Equal(t, expected, string(data))

	err = os.RemoveAll("test")
	assert.Nil(t, err)
}

func TestAddAddsFilesToExistingPlaylist(t *testing.T) {
	playlist := "./test/playlist.txt"
	os.Mkdir("test", os.ModePerm)
	createFiles("test", "dummy1.mp3", "dummy2.mp3", "dummy3.mp3")

	playlistFile, err := os.Create(playlist)
	assert.Nil(t, err)
	defer playlistFile.Close()
	playlistFile.WriteString("dummy1.mp3\n")

	var buffer bytes.Buffer
	err = Run(&buffer, "add", []string{"test"})
	assert.Nil(t, err)

	data, err := os.ReadFile(playlist)
	assert.Nil(t, err)

	expected := "dummy1.mp3\n" +
		"dummy2.mp3\n" +
		"dummy3.mp3\n"

	assert.Equal(t, expected, string(data))

	err = os.RemoveAll("test")
	assert.Nil(t, err)
}

func TestAddOnlyAddsSupportedFiles(t *testing.T) {
	playlist := "./test/playlist.txt"
	os.Mkdir("test", os.ModePerm)
	createFiles("test", "dummy1.ogg", "dummy2.wav", "dummy3.mp3", "dummy.txt", "dummy.mp4")

	var buffer bytes.Buffer
	err := Run(&buffer, "add", []string{"test"})
	assert.Nil(t, err)

	data, err := os.ReadFile(playlist)
	assert.Nil(t, err)

	expected := "dummy1.ogg\n" +
		"dummy2.wav\n" +
		"dummy3.mp3\n"

	assert.Equal(t, expected, string(data))

	err = os.RemoveAll("test")
	assert.Nil(t, err)
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
