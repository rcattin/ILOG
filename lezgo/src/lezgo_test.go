package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// Creating a test tree
func setupTreeTest(tb testing.TB) func(tb testing.TB) {
	os.Mkdir("test", 0755)
	os.Mkdir("test/a", 0755)
	os.Mkdir("test/b", 0755)
	os.Mkdir("test/b/d", 0755)
	os.Mkdir("test/b/e", 0755)
	os.Mkdir("test/b/e/e", 0755)
	os.Mkdir("test/c", 0755)
	os.Mkdir("test/c/d", 0755)
	os.Mkdir("test/c/f", 0755)
	os.Mkdir("test/c/f/g", 0755)
	createFile("test/a/file1.txt", 100)

	return func(tb testing.TB) {
		os.RemoveAll("test")
	}
}

func createFile(n string, s int64) {
	os.Create("test/a/file1.txt")
	os.Truncate("test/a/file1.txt", 100)
}

func TestNoCommand(t *testing.T) {
	out, err := exec.Command("./lezgo").Output()
	if err != nil {
		t.Fail()
	}
	if string(out) != "Expected a command\nList of available commands : search replicates\n" {
		t.Fail()
	}
}

func TestSearchNoArg(t *testing.T) {
	out, err := exec.Command("./lezgo", "search").Output()
	if err != nil {
		t.Fail()
	}
	if string(out) != "Expected an argument\n" {
		t.Fail()
	}
}

func TestSearchA(t *testing.T) {
	teardownFunc := setupTreeTest(t)
	defer teardownFunc(t)

	out, err := exec.Command("./lezgo", "search", "a", "-path=test").Output()
	if err != nil {
		t.Fail()
	}
	wd, _ := os.Getwd()

	if string(out) != wd+"/test/a\n\tSize : 100 octets\n" {
		fmt.Println(string(out))
		t.Fail()
	}
}
