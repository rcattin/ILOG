package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type size struct {
}

func (s *size) visitNode(n *node, d string) bool {
	var size int64 = 0

	readSizeFunc := func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fileInfo, _ := d.Info()
			size += fileInfo.Size()
		}
		return err
	}

	err := filepath.WalkDir(n.path, readSizeFunc)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tSize :", string("\033[31m"), size, "octets", string("\033[0m"))

	return true
}
