package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type search struct {
}

func (s *search) visitNode(n *node, d string) bool {
	//if vFlag {
	// fmt.Println("Searching at", n.path, "...")
	//}

	// Node is found
	if d == filepath.Base(n.path) {
		fmt.Println(n.path, "...")
		// TODO : new thread
		n.accept(&size{}, d)
		return true
	}

	// Node isn't found yet : the search goes on...
	// Get all sub-files
	files, err := ioutil.ReadDir(n.path)
	if err != nil {
		panic(err)
	}

	var foundSubDir bool = false
	// Call search on all directories
	for _, file := range files {
		if file.IsDir() {
			search := &search{}
			newPath := n.path + "/" + file.Name()
			newNode := &node{path: newPath}

			foundSubDir = newNode.accept(search, d) || foundSubDir
		}
	}
	return foundSubDir
}
