package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if a command was given
	if len(os.Args[1:]) == 0 {
		noCmdError()
	}

	// Call the given command
	switch os.Args[1] {
	case "search":
		searchCmd()
		os.Exit(0)

	case "replicates":
		replicatesCmd()
		os.Exit(0)

	default:
		noCmdError()
	}
}

func searchCmd() {

	if len(os.Args[2:]) == 0 {
		fmt.Println("Expected an argument")
		os.Exit(0)
	}

	// Get directory to search
	dFlag := os.Args[2]

	// Get flags
	searchSet := flag.NewFlagSet("search", flag.ExitOnError)
	pFlag := searchSet.String("path", ".", "The directory where the search begins")
	vFlag := searchSet.Bool("v", false /* TODO */, "Explain")
	//hFlag := searchSet.Bool("h", false , "Print help")
	searchSet.Parse(os.Args[3:])

	// Rebuild absolute path of given directory (working directory by default)
	var initPath string
	if !filepath.IsAbs(*pFlag) {
		wd, err := filepath.Abs(".")
		if err != nil {
			panic(err)
		}
		initPath = filepath.Clean(wd + "/" + *pFlag)

	} else {
		initPath = filepath.Clean(*pFlag)
	}

	// TODO : check if directory exists

	// Setting up the search
	initNode := &node{path: initPath}
	search := &search{}
	if *vFlag {
		fmt.Println("Initial path :", initNode.path)
	}

	// Start of the search
	if !initNode.accept(search, dFlag) {
		fmt.Println("Not Found !")
	}
}

func replicatesCmd() {
	fmt.Println("TODO")
}

func noCmdError() {
	fmt.Println("Expected a command")
	fmt.Println("List of available commands : search replicates")
	os.Exit(0)
}
