package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getFileList(folder string, prefix string) []string {
	var files []string

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, prefix) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return files
}

// IOReadDir Returns all the files in the "root" folder that start with "prefix"
func IOReadDir(root string, prefix string) []os.FileInfo {
	var files []os.FileInfo
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileInfo {
		if strings.HasPrefix(file.Name(), prefix) {
			files = append(files, file)
		}
	}

	return files
}
