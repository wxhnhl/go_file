package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	elements := make(map[string]int)
	visit(elements, doc)
	for e, c := range elements {
		fmt.Printf("%s\t%d\n", e, c)
	}
}
func visit(elements map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		visit(elements, i)
	}
}
