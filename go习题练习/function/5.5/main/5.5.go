package main

import (
	"bufio"
	"golang.org/x/net/html"
	"strings"
)

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	return words, images
}
