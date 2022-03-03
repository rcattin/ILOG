package main

type runner interface {
	Init([]string) error
	Run() error
	Name() string
	Help()
}
