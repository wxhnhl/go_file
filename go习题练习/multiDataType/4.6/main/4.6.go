package main

import (
	"fmt"
	"unicode"
)

func main() {

	s := a([]byte{'1', '2', ' ', ' ', '3', ' ', ' ', ' '})
	fmt.Printf("%c %d", s, len(s))

}

func a(s []byte) []byte {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) {
			if unicode.IsSpace(rune(s[i+1])) {
				copy(s[i+1:], s[i+2:])
				s = s[:len(s)-1]
				i--
			}
		}
	}
	return s
}
