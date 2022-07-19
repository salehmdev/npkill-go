package main

import (
	"fmt"
	"log"
	"os"
)

func searchDirs(path string) (bool, error) {
	dir, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f == "node_modules" {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	// scan all directories
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	found, err := searchDirs(path)
	if err != nil {
		log.Fatal(err)
	}

	if found {
		fmt.Println("found!")
	} else {
		fmt.Println("not found...")
	}
	// for 'node_modules', remove all content/folders within that folder.
}
