package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters 不同字符的计数，类型为 map
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings 字符编码长度的计数，类型为 [5]int
	invalid := 0                    // count of invalid UTF-8 characters 无效字符的计数
	var utftype [3]int              //1代表字母 2代表数字 0代表其他类别

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			utftype[1]++
		case unicode.IsNumber(r):
			utftype[2]++
		default:
			utftype[0]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntype\tcount\n")
	for i, n := range utftype {
		var s string
		switch i {
		case 0:
			s = "Other"
		case 1:
			s = "Letter"
		case 2:
			s = "Number"
		}
		fmt.Printf("%s\t%d\n", s, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
