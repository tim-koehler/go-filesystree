package filesystree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilesysTree(t *testing.T) {
	fileList := []string{
		"/foo/bar/baz.txt",
		"/foo/bar/a.txt",
		"/foo/bar/x.go",
		"/tmp/a/b.c",
		"/tmp/b/c.c",
		"/tmp/b/d.c"}

	fst := New()
	for _, s := range fileList {
		fst.AddFile(s, nil)
	}

	dirs := fst.FindDirectoriesAtPath("/tmp")
	assert.Equal(t, len(dirs), 2)
	assert.Equal(t, dirs[0].GetName(), "a")
	assert.Equal(t, dirs[1].GetName(), "b")

	files := fst.FindFilesAtPath("/foo/bar")
	assert.Equal(t, len(files), 3)
	assert.Equal(t, files[0].GetName(), "baz.txt")
	assert.Equal(t, files[1].GetName(), "a.txt")
	assert.Equal(t, files[2].GetName(), "x.go")
}
