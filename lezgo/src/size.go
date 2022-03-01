package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type size struct {
}

func (s *size) visitNode(n *node, args argList) bool {
	var size int64 = 0

	readSizeFunc := func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !dir.IsDir() {
			fileInfo, _ := dir.Info()
			size += fileInfo.Size()
		}
		return err
	}

	err := filepath.WalkDir(n.path, readSizeFunc)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tSize :", size, "octets")

	return true
}
