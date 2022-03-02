package main

import (
	"flag"
)

type Command struct {
	fs   *flag.FlagSet
	args argList
	name string
}
