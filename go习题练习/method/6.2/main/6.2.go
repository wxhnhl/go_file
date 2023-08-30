package main

func (s *IntSet) AddALL(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}
