package main

import (
	"fmt"
)

type size struct {
	size int
}

func (s *size) visitNode(n *node, d string) {
    //do stuff...
    fmt.Println("size of ", d, " ...")
    fmt.Print(s.size)
}
