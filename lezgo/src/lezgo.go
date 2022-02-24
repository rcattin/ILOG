package main

import (
	"path/filepath"
	"fmt"
	"flag"
)

func main() {
	dFlag, pFlag := parseFlags()
	initPath, err := filepath.Abs(pFlag)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	initNode := &node{ path: initPath}
	search := &search{}

	fmt.Println(initNode.path)

	initNode.accept(search, dFlag)
}

func parseFlags() (string, string) {
	var d, p string

        flag.StringVar(&p, "path", ".", "The filepath where the search begins")

	flag.Parse()

	d = flag.Args()[0]
	return d, p
}
