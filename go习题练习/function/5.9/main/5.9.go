package main

import "strings"

func expand(s string, f func(string) string) string {
	str := strings.Replace(s, "foo", f("foo"), -1)
	return str
}
