package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitR struct {
	R io.Reader
	N int64
}

func (l *LimitR) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.N {
		p = p[:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitR{r, n}
}

func main() {
	p := make([]byte, 10)
	str := "12345"
	r := strings.NewReader(str)
	l := LimitReader(r, 4)
	l.Read(p)
	fmt.Println(string(p))
}
