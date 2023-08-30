package main

func (s *IntSet) Len() int {
	ret := 0
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; i < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				ret++
			}
		}
	}
	return ret
}
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &= (^(1 << bit))
}
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}
func (s *IntSet) Copy() *IntSet {
	newWords := make([]uint64, len(s.words))
	copy(newWords, s.words)
	return &IntSet{newWords}
}
