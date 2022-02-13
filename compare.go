package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
)

func main() {
	fileFlag := flag.String("filename", "game-list", "List identifier.")

	purgeFlag := flag.Bool("purge", false, "Delete every compare file except for the two most recent ones.")

	flag.Parse()

	var config = getConfig()

	var fileInfos = IOReadDir(config.Config.FileCompareFolder, config.Config.FilePrefix+*fileFlag+"--")

	if len(fileInfos) < 2 {
		fmt.Println("There are not enough files to compare.")
		return
	}

	files := getAllFiles(fileInfos)

	paramOne, paramTwo := filesToCompare(files)

	paramOne = config.Config.FileCompareFolder + paramOne
	paramTwo = config.Config.FileCompareFolder + paramTwo

	c := exec.Command(config.Config.DiffCommand, paramOne, paramTwo)

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	if *purgeFlag {
		purgeFiles(files, config)
	}

	fmt.Println("Comparison finished.")
}

func getAllFiles(fileInfos []os.FileInfo) []string {
	var files []string

	for _, fileInfo := range fileInfos {
		files = append(files, fileInfo.Name())
	}

	sort.Strings(files)

	return files
}

func filesToCompare(files []string) (string, string) {
	var length = len(files)

	return files[length-2], files[length-1]
}

func purgeFiles(files []string, config Configuration) {
	var length = len(files)

	if length > 2 {
		filesToDelete := files[:length-2]

		for _, file := range filesToDelete {
			os.Remove(config.Config.FileCompareFolder + file)
		}

		fmt.Println(fmt.Sprint("Deleted ", len(filesToDelete), " file(s)."))
	}
}
