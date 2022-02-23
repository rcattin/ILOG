package main

import (

)

func main() {
	initNode := &node{ path: ""}
	search := &search{}
	size := &size{ size : 0}

	initNode.accept(search)
	initNode.accept(size)
}
