package main

import (
	"hash"
	"io/ioutil"
)

// struct symbolisant un fichier qui peut être exister en plusieurs exemplaires,
// d'où le fait que paths soit un tableau, plusieurs chemins
type replicate struct {
	size     int
	hashcode hash.Hash
	paths    []string
	test     int
}

// tableau de replicate, on append un nouvel élément à chaque fois qu'un fichier
// unique est rencontré
type replicates struct {
	files []replicate
}

func (s *replicates) visitNode(n *node, args argList) bool {
	files, err := ioutil.ReadDir(n.path)

	if err != nil {
		panic(err)
	}

	var foundSubDir bool = false

	for _, file := range files {
		if file.IsDir() {
			replicates := s

			newPath := n.path + "/" + file.Name()
			newNode := &node{path: newPath}
			foundSubDir = newNode.accept(replicates, args) || foundSubDir
		}

		// si le fichier existe déjà dans la liste replicates, on ajoute son chemin
		// au champ correspondant

		// ...

		// si il n'existe pas, on append sa signature au tableau replicates

		// ...
	}

	return true
}
