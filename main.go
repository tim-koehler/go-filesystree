package main

import (
	"github.com/tim-koehler/go-filesystree/filesystree"
)

func main() {
	slist := []string{
		"/mnt/sdcard/folder1/a/b/file1.file",
		"/mnt/sdcard/folder1/a/b/file2.file",
		"/mnt/sdcard/folder1/a/b/file3.file",
		"/mnt/sdcard/folder1/a/b/file4.file",
		"/mnt/sdcard/folder1/a/b/file5.file",
		"/mnt/sdcard/folder1/e/c/file6.file",
		"/mnt/sdcard/folder2/d/file7.file",
		"/mnt/sdcard/folder2/d/file8.file",
		"/mnt/sdcard/file9.file"}

	fst := filesystree.New()
	for _, s := range slist {
		fst.AddElement(s)
	}

	fst.PrintTree()
}
