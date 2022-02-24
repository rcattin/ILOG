package main

import (
	"fmt"
	"path/filepath"
)

type search struct{
}

func (s *search) visitNode(n *node, d string) {
	// do stuff
	fmt.Println("searching for ", d ," in ", n.path, "...")
	if (d==filepath.Base(n.path)){
		fmt.Println("Found!")
		n.accept(&size{ size:0}, d)
	}
}
