package main

import "fmt"

func main() {
	a := "hello world!"
	fmt.Println(reverse([]byte(a)))
}

func reverse(arr []byte) string {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return string(arr)
}

//func reverse(s []byte) {
//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//		s[i], s[j] = s[j], s[i]
//	}
//}

//// reverse reverses a slice of ints in place.
//func reverse(s []int) {
//    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//        s[i], s[j] = s[j], s[i]
//    }
//}
