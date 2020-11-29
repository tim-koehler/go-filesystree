package filesystree

import (
	"fmt"
	"strings"
)

// Directory structure containing the name, files, subdirectories, ...
type Directory struct {
	name            string
	files           []*File
	directories     []*Directory
	parent          *Directory
	incrementalPath string
}

func (dir *Directory) add(currentPath string, splitPath []string, meta Metadata) {
	if len(splitPath) == 1 {
		dir.files = append(dir.files, &File{
			name:      splitPath[0],
			path:      currentPath,
			metadata:  meta,
			directory: dir,
		})
		return
	}

	newDir := Directory{
		name:            splitPath[0],
		directories:     []*Directory{},
		files:           []*File{},
		parent:          dir,
		incrementalPath: fmt.Sprintf("%s/%s", currentPath, splitPath[0]),
	}

	contains, nextDir := dir.getIfContains(newDir)
	if !contains {
		dir.directories = append(dir.directories, &newDir)
		newDir.add(newDir.incrementalPath, splitPath[1:], meta)
	} else {
		nextDir.add(nextDir.incrementalPath, splitPath[1:], meta)
	}
}

func (dir *Directory) getIfContains(d Directory) (bool, *Directory) {
	for _, containedDir := range dir.directories {
		if containedDir.incrementalPath == d.incrementalPath {
			return true, containedDir
		}
	}
	return false, nil
}

// GetParentDirectory returns the parent directory of this directory.
func (dir *Directory) GetParentDirectory() *Directory {
	return dir.parent
}

// GetSubDirectories return a list of all the directorys contained in that directory (not recursive).
func (dir *Directory) GetSubDirectories() []*Directory {
	return dir.directories
}

// GetFiles returns a list of all the files contained in that directory (not recursive).
func (dir *Directory) GetFiles() []*File {
	return dir.files
}

// GetFullName returns the path + name of the directory.
func (dir *Directory) GetFullName() string {
	return dir.incrementalPath
}

// GetName returns the name of the directory.
func (dir *Directory) GetName() string {
	return dir.name
}

func (dir *Directory) print(increment int, builder *strings.Builder) string {
	for i := 0; i < increment; i++ {
		builder.WriteRune(' ')
	}
	builder.WriteString(fmt.Sprintf("%s\n", dir.GetName()))

	if len(dir.directories) > 0 {
		for _, d := range dir.directories {
			d.print(increment+2, builder)
		}
	}

	for _, f := range dir.files {
		for i := 0; i < increment; i++ {
			builder.WriteRune(' ')
		}
		builder.WriteString(fmt.Sprintf("- %s\n", f.GetName()))
	}
	return builder.String()
}
