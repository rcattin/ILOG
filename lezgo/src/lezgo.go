package main

import (
	"errors"
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
	// Get directory to search
	// Get flags
	searchSet := flag.NewFlagSet("search", flag.ExitOnError)
	d := searchSet.String("d", "", "The directory to search")
	p := searchSet.String("p", ".", "The directory where the search begins (can be relative or absolute)")
	v := searchSet.Bool("v", false /* TODO */, "Explain")
	h := searchSet.Bool("h", false, "Print help")
	searchSet.Parse(os.Args[2:])

	if *d == "" {
		fmt.Println("Must pass an argument")
		os.Exit(0)
	}
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
		dFlag: *d,
	}
}

func NewGreetCommand() *GreetCommand {
	gc := &GreetCommand{
		fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "name", "World", "name of the person to be greeted")

	return gc
}

type GreetCommand struct {
	fs *flag.FlagSet

	name string
}

func (g *GreetCommand) Name() string {
	return g.fs.Name()
}

func (g *GreetCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *GreetCommand) Run() error {
	fmt.Println("Hello", g.name, "!")
	return nil
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		NewGreetCommand(),
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

/*
func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}*/
