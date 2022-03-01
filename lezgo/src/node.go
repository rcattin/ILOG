package main

type node struct {
	path string
}

func (n *node) accept(v visitor, args argList) bool {
	return v.visitNode(n, args)
}
