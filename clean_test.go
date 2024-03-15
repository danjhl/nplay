package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanRemovesPrefixFromPlaylistFiles(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "tmp")
	defer os.Remove(tmpDir)

	createFiles(tmpDir, "1_dummy1.mp3", "2_dummy2.mp3", "3_dummy3.mp3")
	createPlaylist(tmpDir, "dummy1.mp3", "dummy3.mp3")

	var buffer bytes.Buffer
	err := Run(&buffer, "clean", []string{tmpDir})
	assert.Nil(t, err)

	files := getAllFiles(tmpDir)
	assert.Equal(t, []string{"2_dummy2.mp3", "dummy1.mp3", "dummy3.mp3", "playlist.txt"}, files)
}

func getAllFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames
}
