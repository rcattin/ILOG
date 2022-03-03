package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type searchCmd struct {
	name string
	fs   *flag.FlagSet
	args argList
}

func newSearchCommand() *searchCmd {
	cmd := &searchCmd{
		name: "search",
		fs:   flag.NewFlagSet("search", flag.ExitOnError),
	}

	cmd.fs.StringVar(&cmd.args.pFlag, "p", ".", "The directory where the search begins (can be relative or absolute)")
	cmd.fs.BoolVar(&cmd.args.vFlag, "v", false, "To see where the search is made")
	cmd.fs.BoolVar(&cmd.args.hFlag, "h", false, "To print this help")

	return cmd
}

func (c *searchCmd) Name() string {
	return c.fs.Name()
}

func (c *searchCmd) Init(args []string) error {
	if len(args) == 0 || args[0][0:1] == "-" {
		fmt.Println("Must pass an argument")
		os.Exit(0)
	}
	c.args.dFlag = args[0]
	return c.fs.Parse(args[1:])
}

func (c *searchCmd) Run() error {

	if c.args.hFlag {
		c.Help()
		os.Exit(0)
	}

	// Rebuild absolute path of given directory (working directory by default)
	var absPath string
	if !filepath.IsAbs(c.args.pFlag) {
		wd, err := filepath.Abs(".")
		if err != nil {
			panic(err)
		}
		absPath = filepath.Clean(wd + "/" + c.args.pFlag)

	} else {
		absPath = filepath.Clean(c.args.pFlag)
	}

	// TODO : check if directory exists
	_, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		fmt.Println("Given path does not exist")
		os.Exit(0)
	}

	// Setting up the search
	initNode := &node{path: absPath}
	search := &search{}
	if c.args.vFlag {
		fmt.Println("Initial path :", initNode.path)
	}

	// Start of the search
	if !initNode.accept(search, c.args) {
		fmt.Println("Not Found !")
	}

	return nil
}

func (c *searchCmd) Help() {
	fmt.Println("usage: lezgo search [dirName] [-OPTIONS]")
	fmt.Println()
	fmt.Println("'lezgo search' finds every directory matching the given name in the working directory and prints their size")
	fmt.Println()
	fmt.Println("The flags handled by 'lezgo search' are:")
	c.fs.PrintDefaults()
}
