package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type argList struct {
	pFlag, dFlag string
	vFlag, hFlag bool
}

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
	args := getArgs()

	// TODO : check if directory exists

	// Setting up the search
	initNode := &node{path: args.pFlag}
	search := &search{}
	if args.vFlag {
		fmt.Println("Initial path :", initNode.path)
	}

	// Start of the search
	if !initNode.accept(search, args) {
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

func getArgs() argList {
	if len(os.Args[2:]) == 0 {
		fmt.Println("Expected an argument")
		os.Exit(0)
	}

	// Get directory to search
	d := os.Args[2]
	// Get flags
	searchSet := flag.NewFlagSet("search", flag.ExitOnError)
	p := searchSet.String("path", ".", "The directory where the search begins")
	v := searchSet.Bool("v", false /* TODO */, "Explain")
	h := searchSet.Bool("h", false, "Print help")
	searchSet.Parse(os.Args[3:])

	// Rebuild absolute path of given directory (working directory by default)
	var absPath string
	if !filepath.IsAbs(*p) {
		wd, err := filepath.Abs(".")
		if err != nil {
			panic(err)
		}
		absPath = filepath.Clean(wd + "/" + *p)

	} else {
		absPath = filepath.Clean(*p)
	}

	return argList{
		pFlag: absPath,
		vFlag: *v,
		hFlag: *h,
		dFlag: d,
	}
}
