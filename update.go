package main

import (
	"errors"
	"io/fs"
	"os"
	"regexp"
	"strconv"
)

type Update struct{}

func (u Update) Name() string {
	return "update"
}

func (u Update) Help() string {
	return "update\n\n" +

		"Usage: update [dir]\n\n" +

		"Updates music files inside of a directory by adding numeric prefixes \n" +
		"according to their order in a playlist.txt file. \n" +
		"Files not included in the playlist.txt will not be prefixed.\n\n" +

		"dir: Directory to update (default is current directory)"
}

func (u Update) Execute(args []string) error {
	dir, err := getExecutionDirectory(args)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	entries, err := getPlaylistEntries(dir + "/playlist.txt")
	if err != nil {
		return err
	}

	for i, entry := range entries {
		if entry == "" {
			continue
		}

		var oldFile string

		_, err := os.Stat(dir + "/" + entry)
		if errors.Is(err, os.ErrNotExist) {
			prefixed := findPrefixedFile(entry, &files)
			if prefixed == "" {
				continue
			}
			oldFile = prefixed	
		} else {
			oldFile = entry
		}

		prefix := strconv.Itoa(i+1) + "_"
		err = renameFile(dir, oldFile, prefix + entry)
		if err != nil {
			return err
		}

	}
	return nil
}

func findPrefixedFile(filename string, files *[]fs.DirEntry) string {
	for _, entry := range *files {
		match, _ := regexp.MatchString("\\d+_" + filename, entry.Name()) 
		if match {
			return entry.Name()
		}
	}
	return ""
}

func renameFile(dir string, oldName string, newName string) error {
	return os.Rename(dir + "/" + oldName, dir + "/" + newName)
}
