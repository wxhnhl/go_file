package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var num string
	fmt.Print("please input a number:")
	_, _ = fmt.Scanln(&num)
	fmt.Println(comma_1(num))
}

func comma(s string) string {
	l := len(s)
	var buf bytes.Buffer
	for i := 0; i < l%3; i++ {
		buf.WriteByte(s[i])
	}
	if l%3 != 0 {
		buf.WriteByte(',')
	}
	for i := l % 3; i < l; {
		buf.WriteString(s[i : i+3])
		i += 3
		if i != l {
			buf.WriteByte(',')
		}
	}
	// fmt.Println(buf.String())
	return buf.String()
}

func comma_1(num string) string {
	index := strings.LastIndex(num, ".")
	if index == -1 {
		return comma(num)
	} else {
		return comma(num[:index]) + "." + num[index+1:]
	}
}
