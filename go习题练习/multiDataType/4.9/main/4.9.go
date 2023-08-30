package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq() {
	counts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}
	for k, v := range counts {
		fmt.Println("%s %d\n", k, v)
	}
}
