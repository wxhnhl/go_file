package main

import (
	"fmt"
	"gopl.io/ch2/tempconv"
	"log"
	"os"
	"strconv"
)

func main() {
	for _, str := range os.Args[1:] {
		t, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Fatal(err)
		}
		m := tempconv.Meter(t)
		km := tempconv.Kilometer(t)

		fmt.Printf("%s = %s, %s = %s\n", m, tempconv.MToK(m), km, tempconv.KToM(km))
	}
}
