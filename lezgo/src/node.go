
package main

type node struct{
	path string;
}

func (n *node) accept(v visitor){
	v.visitNode(n)
}
