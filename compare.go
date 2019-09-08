package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
)

func main() {
	var config = getConfig()

	var fileInfos = IOReadDir(config.Config.FileCompareFolder, config.Config.FilePrefix)

	if len(fileInfos) < 2 {
		fmt.Println("There are not enough files to compare.")
		return
	}

	paramOne, paramTwo := filesToCompare(fileInfos)

	paramOne = config.Config.FileCompareFolder + paramOne
	paramTwo = config.Config.FileCompareFolder + paramTwo

	c := exec.Command(config.Config.DiffCommand, paramOne, paramTwo)

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Comparison finished.")
}

func filesToCompare(fileInfos []os.FileInfo) (string, string) {
	var files []string

	for _, fileInfo := range fileInfos {
		files = append(files, fileInfo.Name())
	}

	sort.Strings(files)

	var length = len(files)

	return files[length-2], files[length-1]
}
