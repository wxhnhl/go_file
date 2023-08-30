package main

func rzero(a int) (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = a
		}
	}()
	panic(a)
}
