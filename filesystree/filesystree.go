package filesystree

import (
	"strings"
)

type FilesysTree struct {
	Root *Directory
}

func New() FilesysTree {
	fst := FilesysTree{
		Root: &Directory{
			name:            "/",
			files:           []*File{},
			directories:     []*Directory{},
			incrementalPath: "",
		},
	}
	fst.Root.parent = fst.Root
	return fst
}

func (fst *FilesysTree) AddFile(filePath string, meta *Metadata) {
	if strings.HasPrefix(filePath, "/") {
		filePath = strings.TrimPrefix(filePath, "/")
	}
	splitPath := strings.Split(filePath, "/")
	fst.Root.Add(fst.Root.incrementalPath, splitPath, meta)
}

func (fst *FilesysTree) PrintTree() {
	fst.Root.Print(0)
}

func (fst *FilesysTree) FindDirectoriesAtPath(path string) []*Directory {
	return fst.getDirsAtPath(fst.Root, path)
}

func (fst *FilesysTree) getDirsAtPath(dir *Directory, path string) []*Directory {
	if dir.incrementalPath == path || path == "/" {
		return dir.GetSubDirectories()
	} else if len(dir.GetSubDirectories()) > 0 {
		var result []*Directory
		for _, d := range dir.GetSubDirectories() {
			result = fst.getDirsAtPath(d, path)
			if result != nil {
				return result
			}
		}
	}
	return nil
}

func (fst *FilesysTree) FindFilesAtPath(path string) []*File {
	return fst.getFilesAtPath(fst.Root, path)
}

func (fst *FilesysTree) getFilesAtPath(dir *Directory, path string) []*File {
	if dir.incrementalPath == path || path == "/" {
		return dir.GetFiles()
	} else if len(dir.GetSubDirectories()) > 0 {
		var result []*File
		for _, d := range dir.GetSubDirectories() {
			result = fst.getFilesAtPath(d, path)
			if result != nil {
				return result
			}
		}
	}
	return nil
}
