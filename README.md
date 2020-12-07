# Go-FilesysTree ![Test Badge](https://github.com/tim-koehler/go-filesystree/workflows/Tests/badge.svg) [![Coverage Status](https://coveralls.io/repos/github/tim-koehler/go-filesystree/badge.svg?branch=master)](https://coveralls.io/github/tim-koehler/go-filesystree?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/tim-koehler/go-filesystree)](https://goreportcard.com/report/github.com/tim-koehler/go-filesystree)
  
Go-FilesysTree is a package which I build for another project but decided to publish to maybe help others with building a file system like
tree data structure in Golang.

## Getting Started

```bash
go get github.com/tim-koehler/go-filesystree
```

### Examples

```go
func main() {
    fileList := []string{
        "/foo/bar/baz.txt",
        "/foo/bar/a.txt",
        "/foo/bar/x.go",
        "/tmp/a/b.c",
        "/tmp/b/c.c",
        "/tmp/b/d.c"}

    fst := New()
    for _, s := range fileList {
        fst.AddFile(s, Metadata{
            "Date": time.Now().String(),
        })
    }
}
```

...🚧 work in progress 🚧... 
