package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Print("please input s1:")
	_, _ = fmt.Scanln(&s1)
	fmt.Print("please input s2:")
	_, _ = fmt.Scanln(&s2)
	if order(s1, s2) {
		fmt.Println("IS REVERSED")
	} else {
		fmt.Printf("IS NOT REVERSED")
	}
}

func order(s1, s2 string) bool {
	m := make(map[rune]int)
	for _, i := range s1 {
		m[i]++
	}
	for _, i := range s2 {
		m[i]--
	}
	for _, j := range m {
		if j != 0 {
			return false
		}
	}
	return true
}
