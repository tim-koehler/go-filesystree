package filesystree

import (
	"fmt"
	"strings"
)

type File struct {
	name      string
	path      string
	directory *Directory
	metadata  *Metadata
}

type Metadata struct {
	Date string
	Size int64
}

func (file *File) GetName() string {
	return file.name
}

func (file *File) GetFullName() string {
	return fmt.Sprintf("%s/%s", file.path, file.name)
}

func (file *File) GetDirectory() *Directory {
	return file.directory
}

func (file *File) GetDate() string {
	return strings.Split(file.metadata.Date, " ")[0]
}

func (file *File) GetDateTime() string {
	return file.metadata.Date
}

func (file *File) GetHumanReadableSize() string {
	unit := int64(1000)

	if file.metadata.Size < unit {
		return fmt.Sprintf("%d B", file.metadata.Size)
	}
	div, exp := int64(unit), 0
	for n := file.metadata.Size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(file.metadata.Size)/float64(div), "kMGTPE"[exp])
}
