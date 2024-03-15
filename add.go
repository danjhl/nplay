package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var PLAYLIST string = "playlist.txt"

type Add struct {
}

func (i Add) Name() string {
	return "add"
}

func (i Add) Help() string {
	return "add\n\n" +
		"Usage: add [dir]\n\n" +

		"Add music to a playlist.txt file.\n\n" +

		"dir: optional directory to add files to a playlist.txt (default is current directory)"
}

func (i Add) Execute(args []string) error {
	dir, err := getExecutionDirectory(args)
	if err != nil {
		return err
	}
	playlist := dir + "/" + PLAYLIST

	entries, err := getPlaylistEntries(playlist)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	newEntries := []string{}

	for _, file := range files {
		exists := false
		if !isSupportedFile(file.Name()) {
			continue
		}

		for _, entry := range entries {
			if entry == file.Name() {
				exists = true
				continue
			}
			match, err := regexp.MatchString("\\d+_"+file.Name(), entry)
			if err != nil {
				return err
			}
			if match {
				exists = true
				break
			}
		}

		if !exists {
			newEntries = append(newEntries, file.Name())
		}
	}

	writeEntries := append(entries, newEntries...)

	playlistFile, err := getPlaylistFile(playlist)
	if err != nil {
		return err
	}
	defer playlistFile.Close()
	for _, entry := range writeEntries {
		if entry != "" {
			_, err := playlistFile.WriteString(entry + "\n")
			if err != nil {
				return nil
			}
		}
	}

	return nil
}

func getPlaylistEntries(playlist string) ([]string, error) {
	_, err := os.Stat(playlist)

	if errors.Is(err, os.ErrNotExist) {
		return []string{}, nil
	} else {
		content, err := os.ReadFile(playlist)
		if err != nil {
			return nil, err
		}

		str := strings.ReplaceAll(string(content), "\r\n", "\n")
		entries := strings.Split(str, "\n")
		return entries, nil
	}
}

func getExecutionDirectory(args []string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if len(args) > 0 {
		dir = args[0]
		_, err := os.Stat(dir)
		if errors.Is(err, os.ErrNotExist) {
			return "", errors.New(fmt.Sprintf("Directory '%s' not found", dir))
		}
	}
	return dir, nil
}

func getPlaylistFile(playlist string) (*os.File, error) {
	_, err := os.Stat(playlist)

	if errors.Is(err, os.ErrNotExist) {
		return os.Create(playlist)
	} else {
		return os.OpenFile(playlist, os.O_RDWR|os.O_TRUNC, 0660)
	}
}

func isSupportedFile(fileName string) bool {
	for _, ext := range []string{".mp3", ".wav", ".ogg"} {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
		if strings.HasSuffix(fileName, strings.ToUpper(ext)) {
			return true
		}
	}
	return false

}
