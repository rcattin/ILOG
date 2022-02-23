package main

import (
	"fmt"
)

type size struct {
	size int
}

func (s *size) visitNode(n *node) {
    //do stuff...
    fmt.Println("size of node ...")
    fmt.Print(s.size)
}
