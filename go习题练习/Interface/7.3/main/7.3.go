package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var buf bytes.Buffer
	queue := make([]*tree, 0)
	if t != nil {
		queue = append(queue, t)
	}
	//使用 bfs 遍历树
	buf.WriteByte('[')
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			buf.WriteString(strconv.Itoa(node.value))

			//不在最后一个值后面加空格
			if len(queue) == size && i == size-1 {
				continue
			}
			buf.WriteByte(' ')
		}
		queue = queue[size:]
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	t := new(tree)
	t.value = 1
	t.left = &tree{value: 2}
	t.right = &tree{value: 3}
	fmt.Println(t)
}
