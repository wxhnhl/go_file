package main

import (
	"fmt"
	"io"
)

type S struct {
	str     string
	current int
}

func (s *S) Read(p []byte) (n int, err error) {
	if s.current > len(s.str) {
		return 0, io.EOF
	}
	n = copy(p, s.str)
	s.current = n
	return n, nil
}
func newReader(s string) io.Reader {
	return &S{s, 0}
}

func main() {
	r := newReader("123")
	p := make([]byte, 3)
	r.Read(p)
	fmt.Println(string(p))
}
