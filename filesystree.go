package filesystree

import (
	"strings"
)

// FilesysTree containing the root of the tree.
type FilesysTree struct {
	Root *Directory
}

// New creates a new FilesysTree
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

// AddFile add the file to the filesystem tree.
func (fst *FilesysTree) AddFile(filePath string, meta Metadata) {
	if strings.HasPrefix(filePath, "/") {
		filePath = strings.TrimPrefix(filePath, "/")
	}
	splitPath := strings.Split(filePath, "/")
	fst.Root.add(fst.Root.incrementalPath, splitPath, meta)
}

// FindDirectoriesAtPath return a list of directories at a given path.
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

// FindFilesAtPath return a list of files at a given path.
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

// GetTree return the whole filesystem in a tree like string pattern.
func (fst *FilesysTree) GetTree() string {
	builder := strings.Builder{}
	return fst.Root.print(0, &builder)
}
