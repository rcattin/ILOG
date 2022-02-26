package main

type node struct {
	path string
}

func (n *node) accept(v visitor, d string) bool {
	return v.visitNode(n, d)
}
