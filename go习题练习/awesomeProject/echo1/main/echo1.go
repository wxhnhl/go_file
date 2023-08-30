package main

import (
	"fmt"
	"os"
)

func main() {

	for line, arg := range os.Args[0:] {
		fmt.Println("num =", line, ", arg = ", arg)
	}
}
