package main

import (
	"path/filepath"
	"fmt"
	"flag"
	"os"
)

func main() {
	dFlag, pFlag := parseFlags()

	if (!filepath.IsAbs(pFlag)){
		wd, err := filepath.Abs(".")
		if err != nil {
		        panic(err)
		}
		initPath := filepath.Clean(wd + "/" + pFlag)

	}else{
		initPath := filepath.Clean(pFlag)
	}

	// TODO : check if filepath is real

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
