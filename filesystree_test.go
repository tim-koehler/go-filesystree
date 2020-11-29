package filesystree

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFilesysTree(t *testing.T) {
	var dateTime time.Time
	fileList := []string{
		"/foo/bar/baz.txt",
		"/foo/bar/a.txt",
		"/foo/bar/x.go",
		"/tmp/a/b.c",
		"/tmp/b/c.c",
		"/tmp/b/d.c"}

	fst := New()
	for _, s := range fileList {
		dateTime = time.Now()
		fst.AddFile(s, Metadata{
			"Date": dateTime.String(),
		})
	}

	dirs := fst.FindDirectoriesAtPath("/tmp")
	assert.Equal(t, len(dirs), 2)
	assert.Equal(t, dirs[0].GetName(), "a")
	assert.Equal(t, dirs[1].GetName(), "b")

	files := fst.FindFilesAtPath("/foo/bar")
	assert.Equal(t, len(files), 3)
	assert.Equal(t, files[0].GetName(), "baz.txt")
	assert.Equal(t, files[1].GetName(), "a.txt")
	assert.Equal(t, files[1].GetFullName(), "/foo/bar/a.txt")
	assert.Equal(t, files[2].GetName(), "x.go")
	assert.Equal(t, files[2].GetDirectory().GetName(), "bar")
	assert.Equal(t, files[2].GetDirectory().GetFullName(), "/foo/bar")
	assert.Equal(t, files[2].GetDirectory().GetParentDirectory().GetFullName(), "/foo")

	files = fst.FindFilesAtPath("/tmp/b")
	assert.Equal(t, files[1].GetMetadata()["Date"], dateTime.String())
	assert.Equal(t,
		base64.StdEncoding.EncodeToString([]byte(fst.GetTree())),
		"LwogIGZvbwogICAgYmFyCiAgICAtIGJhei50eHQKICAgIC0gYS50eHQKICAgIC0geC5nbwogIHRtcAogICAgYQogICAgLSBiLmMKICAgIGIKICAgIC0gYy5jCiAgICAtIGQuYwo=")
}
