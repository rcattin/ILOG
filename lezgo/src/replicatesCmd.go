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
	cmd.fs.BoolVar(&cmd.args.vFlag, "v", false, "Extra details")

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
		fmt.Println("No duplicates found!")
	} else {
		fmt.Println("Duplicates found :")
		fmt.Println()
		for _, rep := range replicates.files {
			if rep.numReplicates > 1 {
				fmt.Println(rep.numReplicates, "duplicates of this file : ")
				for _, path := range rep.paths {
					fmt.Print(path)
					if c.args.vFlag {
						fmt.Printf("  \tHashCode : %x", rep.hashcode)
						fmt.Printf("  \tSize : %d bytes", rep.size)
					}
					fmt.Print("\n")
				}
				fmt.Println()
			}
		}
	}

	return nil
}

func (c *replicatesCmd) Help() {
	fmt.Println("usage: lezgo replicates [-OPTIONS]")
	fmt.Println()
	fmt.Println("'lezgo replicates' finds all identical documents in the working directory")
	fmt.Println()
	fmt.Println("The flags handled by 'lezgo replicates' are:")
	c.fs.PrintDefaults()
}
