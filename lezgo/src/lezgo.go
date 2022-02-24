package main

import (
	"path/filepath"
	"fmt"
	"flag"
)

func main() {
	dFlag := parseFlags()

	initPath, err := filepath.Abs(dFlag)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	initNode := &node{ path: initPath}
	search := &search{}
	size := &size{ size : 0}

	fmt.Println(initNode.path)

	initNode.accept(search)
	initNode.accept(size)
}

func parseFlags() string {
	var d string

        flag.StringVar(&d, "d", ".", "The filepath where the search begins")

	flag.Parse()
	return d
}
