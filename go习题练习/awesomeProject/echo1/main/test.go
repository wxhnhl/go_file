package main

import (
	"strings"
)

func stringJoin() string {
	input := []string{"Welcome", "To", "China"}
	result := strings.Join(input, " ")
	return result
}

func stringPlus() string {
	input := []string{"Welcome", "To", "China"}
	var result, sep string
	for _, j := range input {
		result += sep + j
		sep = " "
	}
	return result
}
func main() {

}
