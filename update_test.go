package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateAddsPrefix(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "tmp")
	defer os.Remove(tmpDir)

	createFiles(tmpDir, "dummy1.mp3", "dummy2.mp3", "dummy3.mp3")
	createPlaylist(tmpDir, "dummy3.mp3", "dummy1.mp3", "dummy2.mp3")

	var buffer bytes.Buffer
	err := Run(&buffer, "update", []string{tmpDir})
	assert.Nil(t, err)

	fileNames, err := getSupportedFiles(tmpDir)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1_dummy3.mp3", "2_dummy1.mp3", "3_dummy2.mp3"}, fileNames)
}

func TestUpdateAddsPrefixToNewEntry(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "tmp")
	defer os.Remove(tmpDir)

	createFiles(tmpDir, "1_dummy1.mp3", "2_dummy2.mp3", "dummy3.mp3")
	createPlaylist(tmpDir, "dummy3.mp3", "dummy1.mp3", "dummy2.mp3")

	var buffer bytes.Buffer
	err := Run(&buffer, "update", []string{tmpDir})
	assert.Nil(t, err)

	fileNames, err := getSupportedFiles(tmpDir)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1_dummy3.mp3", "2_dummy1.mp3", "3_dummy2.mp3"}, fileNames)
}

func TestUpdateWithFileNotOnPlaylist(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "tmp")
	defer os.Remove(tmpDir)

	createFiles(tmpDir, "dummy1.mp3", "dummy2.mp3", "dummy3.mp3")
	createPlaylist(tmpDir, "dummy3.mp3", "dummy2.mp3")

	var buffer bytes.Buffer
	err := Run(&buffer, "update", []string{tmpDir})
	assert.Nil(t, err)

	fileNames, err := getSupportedFiles(tmpDir)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1_dummy3.mp3", "2_dummy2.mp3", "dummy1.mp3"}, fileNames)
}

func getSupportedFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	fileNames := []string{}
	for _, file := range files {
		if !isSupportedFile(file.Name()) {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}
