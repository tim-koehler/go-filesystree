package filesystree

import (
	"strings"
)

type FilesysTree struct {
	Root *Directory
}

func New() FilesysTree {
	return FilesysTree{
		Root: &Directory{
			Name:            "root",
			Files:           []*File{},
			Directories:     []*Directory{},
			incrementalPath: "root",
		},
	}
}

func (fst *FilesysTree) AddElement(filePath string) {
	if strings.HasPrefix(filePath, "/") {
		filePath = strings.TrimPrefix(filePath, "/")
	}
	splitPath := strings.Split(filePath, "/")
	fst.Root.Add(fst.Root.incrementalPath, splitPath)
}

func (fst *FilesysTree) PrintTree() {
	fst.Root.Print(0)
}
