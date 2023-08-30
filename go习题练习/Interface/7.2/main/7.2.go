package main

import (
	"io"
)

var bytes int64

type Counts struct {
	w io.Writer
}

func (c *Counts) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	bytes += int64(n)
	return
}

func main() {

}
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	return &Counts{w}, &bytes
}
