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
	os.Mkdir("test/b/i", 0755)
	os.Mkdir("test/b/e", 0755)
	os.Mkdir("test/b/e/e", 0755)
	os.Mkdir("test/c", 0755)
	os.Mkdir("test/c/d", 0755)
	os.Mkdir("test/c/f", 0755)
	os.Mkdir("test/c/f/g", 0755)
	os.Mkdir("test/h", 0755)
	createFile("test/a/file100.txt", 100)
	createFile("test/h/file10.txt", 10)
	createFile("test/h/file25.txt", 25)
	createFile("test/b/i/file25.txt", 25)
	createFile("test/h/file31.txt", 31)
	createFile("test/c/f/file50.txt", 50)
	createFile("test/c/f/g/file100.txt", 100)

	return func(tb testing.TB) {
		os.RemoveAll("test")
	}
}

func createFile(n string, s int64) {
	os.Create(n)
	os.Truncate(n, s)
}

func TestNoCommand(t *testing.T) {
	out, err := exec.Command("./lezgo").Output()
	if err != nil {
		t.Fatalf(err.Error())
		t.Fail()
	}
	want := "Expected a command\nList of available commands : search replicates\n"
	if string(out) != want {
		t.Fatalf("\ngot \n%vwant\n%v", string(out), want)
		t.Fail()
	}
}

func TestBadCommand(t *testing.T) {
	out, err := exec.Command("./lezgo", "test").Output()
	if err != nil {
		t.Fatalf(err.Error())
		t.Fail()
	}
	want := "Unknown subcommand: test\n"
	if string(out) != want {
		t.Fatalf("\ngot \n%vwant\n%v", out, []byte(want))
		t.Fail()
	}
}

func TestSearchNoArg(t *testing.T) {
	out, err := exec.Command("./lezgo", "search").Output()
	if err != nil {
		t.Fail()
	}
	want := "Must pass an argument\n"
	if string(out) != want {
		t.Errorf("\ngot \n%vwant\n%v", string(out), want)
		t.Fail()
	}
}

func TestSearchBadPath(t *testing.T) {
	out, err := exec.Command("./lezgo", "search", "-d=test", "-p=not/a/path").Output()
	if err != nil {
		t.Fail()
	}
	want := "Given path does not exist\n"
	if string(out) != want {
		t.Errorf("\ngot \n%vwant\n%v", string(out), want)
		t.Fail()
	}
}

func TestSearch(t *testing.T) {
	// Getting working directory
	wd, _ := os.Getwd()

	// Test cases
	cases := []struct {
		name  string
		input []string
		want  string
	}{
		{
			"OneDir-OneFile",
			[]string{"search", "-d=a", "-p=test"},
			wd + "/test/a\n\tSize : 100 octets\n",
		},
		{
			"OneDir-ThreeFiles",
			[]string{"search", "-d=h", "-p=test"},
			wd + "/test/h\n\tSize : 66 octets\n",
		},
		{
			"TwoDir-Empty",
			[]string{"search", "-d=d", "-p=test"},
			wd + "/test/b/d\n\tSize : 0 octets\n" + wd + "/test/c/d\n\tSize : 0 octets\n",
		},
		{
			"OneDir-SubDirAndFile",
			[]string{"search", "-d=f", "-p=test"},
			wd + "/test/c/f\n\tSize : 150 octets\n",
		},
		// Add tests here
	}

	// Running tests
	for i, tc := range cases {
		t.Run(fmt.Sprintf("Test %d : %s", i+1, tc.name), func(t *testing.T) {
			// setUp + tearDown
			teardownFunc := setupTreeTest(t)
			defer teardownFunc(t)

			got, err := exec.Command("./lezgo", tc.input...).Output()
			if err != nil {
				t.Fail()
			}

			if string(got) != tc.want {
				t.Fatalf("\ngot \n%vwant\n%v", string(got), tc.want)
				t.Fail()
			}
		})
	}
}
