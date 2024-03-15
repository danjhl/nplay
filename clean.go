package main

import (
	"io/ioutil"
	"os"
	"regexp"
)

type Clean struct{}

func (c Clean) Name() string {
	return "clean"
}

func (c Clean) Help() string {
	return "clean\n\n" +
		"usage: clean [dir]\n\n" +
		"Removes prefix from filenames in playlist.txt\n\n" +
		"dir: Directory to clean (default is current directory)"
}

func (c Clean) Execute(args []string) error {
	dir, err := getExecutionDirectory(args)
	if err != nil {
		return err
	}

	entries, err := getPlaylistEntries(dir + "/playlist.txt")
	if err != nil {
		return err
	}

	fileNames, err := getMusicFiles(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		r := regexp.MustCompile("(?P<prefix>\\d+_)" + regexp.QuoteMeta(entry))
		var prefix string

		for _, fileName := range fileNames {
			matches := r.FindStringSubmatch(fileName)
			if len(matches) == 0 {
				continue
			}
			prefix = matches[r.SubexpIndex("prefix")]
			break
		}

		if prefix != "" {
			err = os.Rename(dir+"/"+prefix+entry, dir+"/"+entry)
		}
	}

	return nil
}

func getMusicFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}
