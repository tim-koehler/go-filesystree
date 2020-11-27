package filesystree

import (
	"fmt"
)

type Directory struct {
	Name            string
	Files           []*File
	Directories     []*Directory
	incrementalPath string
}

type File struct {
	Name string
}

func (dir *Directory) Add(currentPath string, splitPath []string) {
	if len(splitPath) == 1 {
		dir.Files = append(dir.Files, &File{Name: splitPath[0]})
		return
	}

	newDir := Directory{
		Name:            splitPath[0],
		Directories:     []*Directory{},
		Files:           []*File{},
		incrementalPath: fmt.Sprintf("%s/%s", currentPath, splitPath[0]),
	}

	contains, nextDir := dir.getIfContains(newDir)
	if !contains {
		dir.Directories = append(dir.Directories, &newDir)
		newDir.Add(newDir.incrementalPath, splitPath[1:])
	} else {
		nextDir.Add(nextDir.incrementalPath, splitPath[1:])
	}
}

func (dir *Directory) Print(increment int) {
	for i := 0; i < increment; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(dir.String())

	if len(dir.Directories) > 0 {
		for _, d := range dir.Directories {
			d.Print(increment + 2)
		}
	}

	for _, f := range dir.Files {
		for i := 0; i < increment; i++ {
			fmt.Printf(" ")
		}
		fmt.Println("- " + f.String())
	}
}

func (dir *Directory) getIfContains(d Directory) (bool, *Directory) {
	for _, containedDir := range dir.Directories {
		if containedDir.incrementalPath == d.incrementalPath {
			return true, containedDir
		}
	}
	return false, nil
}

func (dir *Directory) String() string {
	return dir.Name
}

func (file *File) String() string {
	return file.Name
}
