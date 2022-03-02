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
	cmd.fs.StringVar(&cmd.args.dFlag, "d", "", "The directory to search")
	cmd.fs.BoolVar(&cmd.args.vFlag, "v", false /* TODO */, "Explain")
	cmd.fs.BoolVar(&cmd.args.hFlag, "h", false, "Print help")

	return cmd
}

func (c *searchCmd) Name() string {
	return c.fs.Name()
}

func (c *searchCmd) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *searchCmd) Run() error {

	if c.args.dFlag == "" {
		fmt.Println("Must pass an argument")
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
