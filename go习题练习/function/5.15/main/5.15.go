package main

func max(vals ...int) (m int) {
	if len(vals) == 0 {
		return 0
	}
	if len(vals) == 1 {
		return vals[0]
	}
	m = vals[0]
	for i := 1; i < len(vals); i++ {
		if m > vals[i] {
			m = vals[i]
		}
	}
	return m
}
func min(vals ...int) (m int) {
	if len(vals) == 0 {
		return 0
	}
	if len(vals) == 1 {
		return vals[0]
	}
	m = vals[0]
	for i := 1; i < len(vals); i++ {
		if m < vals[i] {
			m = vals[i]
		}
	}
	return m
}
