package main

type visitor interface {
	visitNode(*node, argList) bool
}
