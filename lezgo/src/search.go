package main

import (
	"fmt"
)

type search struct{
}

func (s *search) visitNode(n *node) {
	// do stuff
	fmt.Println("searching node...")
}
