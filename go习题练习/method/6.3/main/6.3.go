package main

import "math"

func (s *IntSet) And(t *IntSet) {
	min := math.Min(len(s.words), len(t.words))
	for i := 0; i < min; i++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[:min]
}
