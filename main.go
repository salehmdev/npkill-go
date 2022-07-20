package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func findNodeModsCountRecurive(path string) (int, error) {
	dir, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer dir.Close()

	files, err := dir.Readdirnames(0)
	if err != nil {
		return 0, err
	}

	nodeModulesCount := 0

	for _, file := range files {
		if file == "node_modules" {
			nodeModulesCount += 1
			continue
		}

		filePath := path + string(os.PathSeparator) + file
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return 0, err
		}

		if fileInfo.IsDir() {
			newCount, err := findNodeModsCountRecurive(filePath)
			if err != nil {
				return 0, err
			}
			nodeModulesCount += newCount
		}
	}

	return nodeModulesCount, nil
}

func main() {
	// scan all directories
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	foundCount, err := findNodeModsCountRecurive(path)
	if err != nil {
		log.Fatal(err)
	}

	if foundCount > 0 {
		fmt.Println("Found " + strconv.Itoa(foundCount) + " \"node_modules\" directories.")
		os.RemoveAll("node_modules")
	} else {
		fmt.Println("No \"node_modules\" found...")
	}
	// for 'node_modules', remove all content/folders within that folder.
}
