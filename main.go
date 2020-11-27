package main

import (
	"github.com/tim-koehler/go-filesystree/filesystree"
)

func main() {
	slist := []string{
		"/foo/bar/baz.txt",
		"/foo/bar/a.txt",
		"/tmp/a/b.c",
		"/tmp/b/c.c",
		"/tmp/b/d.c"}

	fst := filesystree.New()
	for _, s := range slist {
		fst.AddFile(s, nil)
	}

	fst.PrintTree()
}
