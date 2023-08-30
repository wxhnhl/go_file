package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"算法":  {"数据结构"},
	"微积分": {"线性代数"},
	"编译器": {
		"数据结构",
		"形式语言",
		"计算机组成",
	},
	"数据结构":  {"离散数学"},
	"数据库":   {"数据结构"},
	"离散数学":  {"编程入门"},
	"形式语言":  {"离散数学"},
	"计算机网络": {"操作系统"},
	"操作系统":  {"数据结构", "计算机组成"},
	"编程语言":  {"数据结构", "计算机组成"},
}

func main() {
	for k, v := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", k, v)
	}
}

func topoSort(m map[string][]string) map[int]string {

	order := make(map[int]string)
	seen := make(map[string]bool)
	index := 1
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				//order = append(order, item)
				order[index] = item
				index++
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}
