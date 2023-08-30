package main

//map去重
func nonEqual(strings []string) []string {
	m1 := make(map[string]int)

	for i := 0; i < len(strings); i++ {
		if _, ok := m1[strings[i]]; ok {
			i++
			continue
		} else {
			m1[strings[i]] = i
		}
	}
	i := 0
	for k, _ := range m1 {
		strings[i] = k
		i++
	}
	return strings[:i]
}

//双循环去重
func nonEqual1(strings []string) []string {
	for i := 0; i < len(strings)-1; i++ {
		for j := i + 1; j < len(strings); j++ {
			if strings[i] == strings[j] {
				copy(strings[j:], strings[j+1:])
				strings = strings[:len(strings)-1]
				j--
			}
		}
	}
	return strings
}
