package core

import (
	"io"
	"os"
)

func NewInMemoryFile(contents string) InMemoryFile {
	return InMemoryFile{Contents: []byte(contents)}
}

type InMemoryFile struct {
	Contents []byte
}

func (InMemoryFile) Close() error { return nil }
func (file InMemoryFile) Read(p []byte) (int, error) {
	contents := file.Contents
	end := min(len(p), len(contents))
	copiedBytes := copy(p, contents[0:end])
	if len(p) > len(contents) {
		return copiedBytes, io.EOF
	}
	return copiedBytes, nil
}
func (InMemoryFile) Seek(offset int64, whence int) (int64, error) { panic("not implemented") }
func (InMemoryFile) Readdir(count int) ([]os.FileInfo, error)     { panic("not implemented") }
func (InMemoryFile) Stat() (os.FileInfo, error)                   { panic("not implemented") }

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
