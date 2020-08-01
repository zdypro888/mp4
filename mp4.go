package mp4

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/alfg/mp4/atom"
)

// Open opens a file and returns a &File{}.
func Open(path string) (f *atom.File, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	f = &atom.File{
		Reader: file,
	}

	return f, f.Parse()
}

// OpenFromReader read and returns a &File{}.
func OpenFromReader(reader io.ReaderAt, size int) (f *atom.File, err error) {
	f = &atom.File{
		Reader: reader,
		Size:   int64(size),
	}
	return f, f.Parse()
}

// OpenFromBytes read and returns a &File{}.
func OpenFromBytes(buffer []byte) (f *atom.File, err error) {
	f = &atom.File{
		Reader: bytes.NewReader(buffer),
		Size:   int64(len(buffer)),
	}
	return f, f.Parse()
}
