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

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("Expected a command\nList of available commands : search replicates")
	}

	cmds := []Runner{
		newSearchCommand(),
		newReplicatesCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
