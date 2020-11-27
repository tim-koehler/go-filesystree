package filesystree

import (
	"fmt"
)

type Directory struct {
	name            string
	files           []*File
	directories     []*Directory
	parent          *Directory
	incrementalPath string
}

func (dir *Directory) Add(currentPath string, splitPath []string, meta *Metadata) {
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
		newDir.Add(newDir.incrementalPath, splitPath[1:], meta)
	} else {
		nextDir.Add(nextDir.incrementalPath, splitPath[1:], meta)
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

func (dir *Directory) GetParentDirectory() *Directory {
	return dir.parent
}

func (dir *Directory) GetSubDirectories() []*Directory {
	return dir.directories
}

func (dir *Directory) GetFiles() []*File {
	return dir.files
}

func (dir *Directory) GetAbsolutPath() string {
	return dir.incrementalPath
}

func (dir *Directory) GetName() string {
	return dir.name
}

func (dir *Directory) Print(increment int) {
	for i := 0; i < increment; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(dir.GetName())

	if len(dir.directories) > 0 {
		for _, d := range dir.directories {
			d.Print(increment + 2)
		}
	}

	for _, f := range dir.files {
		for i := 0; i < increment; i++ {
			fmt.Printf(" ")
		}
		fmt.Println("- " + f.GetName())
	}
}
