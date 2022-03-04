package main

import (
	"crypto/sha1"
	"io/ioutil"
	"os"
)

// struct symbolisant un fichier qui peut être exister en plusieurs exemplaires,
// d'où le fait que paths soit un tableau, plusieurs chemins
type replicate struct {
	size          uint64
	hashcode      []byte
	paths         []string
	numReplicates int
}

// tableau de replicate, on append un nouvel élément à chaque fois qu'un fichier
// unique est rencontré
type replicates struct {
	files []replicate
	size  uint
}

func computeHash(file []byte) []byte {
	h := sha1.New()
	h.Write(file)
	r := h.Sum(nil)
	return r
}

func registerFile(s *replicates, fileInfo os.FileInfo, path string) bool {

	replicateFound := false
	for i := 0; i < int(s.size); i++ {
		if s.files[i].size == uint64(fileInfo.Size()) { // fichier de même taille
			if len(s.files[i].hashcode) == 0 { // si son hash n'est pas calculé on le calcule
				file, err := ioutil.ReadFile(s.files[i].paths[0])
				if err != nil {
					panic(err)
				}
				s.files[i].hashcode = computeHash(file)
				//fmt.Printf("%x", s.files[i].hashcode)
			}
			fpath := path + "/" + fileInfo.Name()
			file, err := ioutil.ReadFile(fpath)
			if err != nil {
				panic(err)
			}
			a, b := string(s.files[i].hashcode), string(computeHash(file))
			if a == b {
				replicateFound = true
				s.files[i].paths = append(s.files[i].paths, fpath)
				s.files[i].numReplicates++
			}
		}
	}

	if replicateFound {

	} else {
		r := replicate{uint64(fileInfo.Size()), []byte{}, []string{path + "/" + fileInfo.Name()}, 1}
		s.files = append(s.files, r)
		s.size++
	}

	return replicateFound
}

func (s *replicates) visitNode(n *node, args argList) bool {
	files, err := ioutil.ReadDir(n.path)

	if err != nil {
		panic(err)
	}

	var foundSubDir bool = false
	var replicateFound bool = false
	for _, file := range files {
		if file.IsDir() {
			replicates := s

			newPath := n.path + "/" + file.Name()
			newNode := &node{path: newPath}
			replicateFound = replicateFound || newNode.accept(replicates, args) || foundSubDir
		} else {
			replicateFound = registerFile(s, file, n.path)
		}
	}

	return replicateFound
}
