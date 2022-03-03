package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

type replicatesCmd struct {
	name string
	fs   *flag.FlagSet
	args argList
}

func newReplicatesCommand() *replicatesCmd {
	cmd := &replicatesCmd{
		name: "replicates",
		fs:   flag.NewFlagSet("replicates", flag.ExitOnError),
	}

	cmd.fs.StringVar(&cmd.args.pFlag, "p", ".", "The directory where the search begins (can be relative or absolute)")

	return cmd
}

func (c *replicatesCmd) Name() string {
	return c.fs.Name()
}

func (c *replicatesCmd) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *replicatesCmd) Run() error {

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
	replicates := &replicates{}
	if c.args.vFlag {
		fmt.Println("Initial path :", initNode.path)
	}

	// Start of the search
	if !initNode.accept(replicates, c.args) {
		fmt.Println("Not Found !")
	}

	return nil
}

func (c *replicatesCmd) Help() {}
