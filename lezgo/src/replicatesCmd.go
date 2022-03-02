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

	// rajoute tes flags ici (pas besoin de les Parse())

	return cmd
}

func (c *replicatesCmd) Name() string {
	return c.fs.Name()
}

func (c *replicatesCmd) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *replicatesCmd) Run() error {

	// rajoute ton main ici
	fmt.Println("todo @Eloi")

	return nil
}
