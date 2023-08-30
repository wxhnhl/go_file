package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, n int) []int {
	result := make([]int, len(s))
	index := len(s) - n
	for i := 0; i < len(s); i++ {

		if index >= len(s) {
			index = 0
		}
		result[i] = s[index]
		index++
	}
	return result
}
