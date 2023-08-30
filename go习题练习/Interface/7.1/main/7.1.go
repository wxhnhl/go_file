package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(bytes.NewReader(p))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {

	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

}
