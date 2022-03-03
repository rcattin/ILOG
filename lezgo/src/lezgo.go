package main

import (
	"errors"
	"fmt"
	"os"
)

type argList struct {
	pFlag, dFlag string
	vFlag, hFlag bool
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("Expected a command\nList of available commands : search replicates")
	}

	cmds := []runner{
		newSearchCommand(),
		newReplicatesCommand(),
		newHelpCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			err := cmd.Init(os.Args[2:])
			if err != nil {
				return err
			}

			return cmd.Run()
		}
	}

	return errors.New("Unknown subcommand: " + subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
