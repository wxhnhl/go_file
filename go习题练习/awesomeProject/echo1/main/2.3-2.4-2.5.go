package main

import (
	"fmt"
	"gopl.io/ch2/popcount"
)

func main() {
	fmt.Printf("popcount:%d\n", popcount.PopCount(1000))
	fmt.Printf("2.2:%d\n", popcount.PopCount1(1000))
	fmt.Printf("2.3:%d\n", popcount.PopCount2(1000))
	fmt.Printf("2.4:%d\n", popcount.PopCount3(1000))
}
