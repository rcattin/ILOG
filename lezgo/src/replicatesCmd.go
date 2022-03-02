package main

import (
	"flag"
	"fmt"
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

	return cmd
}

func (c *replicatesCmd) Name() string {
	return c.fs.Name()
}

func (c *replicatesCmd) Init(args []string) error {
	fmt.Println(c.args)
	return c.fs.Parse(args)
}

func (c *replicatesCmd) Run() error {

	fmt.Println("todo @Eloi")

	return nil
}
