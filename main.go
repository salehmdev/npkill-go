package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getNodeModsPathsRecurive(path string) ([]string, error) {
	dir, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer dir.Close()

	files, err := dir.Readdirnames(0)
	if err != nil {
		return []string{}, err
	}

	var nodeModulesPathArr []string

	for _, file := range files {
		if file == "node_modules" {
			nodeModulesPathArr = append(nodeModulesPathArr, path+"/"+file)
			continue
		}

		filePath := path + string(os.PathSeparator) + file
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return []string{}, err
		}

		if fileInfo.IsDir() {
			newNodeModulesPathArr, err := getNodeModsPathsRecurive(filePath)
			if err != nil {
				return []string{}, err
			}
			nodeModulesPathArr = append(nodeModulesPathArr, newNodeModulesPathArr...)
		}
	}

	return nodeModulesPathArr, nil
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	foundArr, err := getNodeModsPathsRecurive(path)
	if err != nil {
		log.Fatal(err)
	}

	if len(foundArr) > 0 {
		fmt.Println("Found " + strconv.Itoa(len(foundArr)) + " \"node_modules\" directories.")
		fmt.Println("Deleting...")
		for _, f := range foundArr {
			fmt.Println(f)
			os.RemoveAll(f)
		}
		fmt.Println("Done! :)")
	} else {
		fmt.Println("No \"node_modules\" found... :)")
	}
}
