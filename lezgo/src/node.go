
package main

type node struct{
	path string;
}

func (n *node) accept(v visitor, d string){
	v.visitNode(n, d)
}
