package main

type visitor interface {
	visitNode(*node, string) bool
}
