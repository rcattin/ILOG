package main

import (
	"flag"
	"fmt"
	"os"
)

type helpCmd struct {
	name string
	fs   *flag.FlagSet
	args argList
}

func newHelpCommand() *helpCmd {
	cmd := &helpCmd{
		name: "help",
		fs:   flag.NewFlagSet("help", flag.ExitOnError),
	}
	return cmd
}

func (c *helpCmd) Name() string {
	return c.fs.Name()
}

func (c *helpCmd) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *helpCmd) Run() error {

	cmds := []runner{
		newSearchCommand(),
		newReplicatesCommand(),
	}

	var helpArg string = ""
	if len(os.Args) >= 3 {
		helpArg = os.Args[2]
	}

	for _, cmd := range cmds {
		if cmd.Name() == helpArg {
			cmd.Help()
			return nil
		}
	}
	c.Help()

	return nil
}

func (c *helpCmd) Help() {
	fmt.Println("lezgo is a small command-line-interface with two sub commands :")
	fmt.Println("\t'lezgo search' finds every directory matching the given name in the working directory and prints their size")
	fmt.Println("\t'lezgo replicates' does stuff")
}
