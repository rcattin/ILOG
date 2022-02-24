package main

import (
	"fmt"
	"path/filepath"
	"io/ioutil"
)

type search struct{
}

func (s *search) visitNode(n *node, d string) {
	fmt.Println("searching for ", d ," in ", n.path, "...")

	// found node
	if (d==filepath.Base(n.path)){
		fmt.Println("Found!")

		// TODO : new thread
		n.accept(&size{ size:0}, d)
		return
	}

	// node not found yet : continue searching
	files, err := ioutil.ReadDir(n.path)
	if err!= nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir(){
			search := &search{}
			newPath := n.path+"/"+file.Name()
			newNode := &node{ path: newPath}
			newNode.accept(search, d)
		}
	}
}
