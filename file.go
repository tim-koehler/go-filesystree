package filesystree

import (
	"fmt"
)

// File struct which contains name, path, directory and optional metadata.
type File struct {
	name      string
	path      string
	directory *Directory
	metadata  Metadata
}

// Metadata can be used to store addidional metadata for files.
type Metadata map[string]interface{}

// GetName returns the name of the file.
func (file *File) GetName() string {
	return file.name
}

// GetFullName returns the name of the path + file.
func (file *File) GetFullName() string {
	return fmt.Sprintf("%s/%s", file.path, file.name)
}

// GetDirectory returns the the directory which contains this file.
func (file *File) GetDirectory() *Directory {
	return file.directory
}

// GetMetadata returns the attached metadata of the file.
// If no metadata is defined nil is returned.
func (file *File) GetMetadata() Metadata {
	return file.metadata
}
