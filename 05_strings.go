package main

type Strings []string

func (s *Strings) Len() int {
	return len(*s)
}

func (s *Strings) Swap(i, j int) []string {
	if s.Less(i, j) {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		return *s
	}
	return *s
}

func (s *Strings) Less(i, j int) bool {
	if (*s)[i] > (*s)[j]{
		return false
	}
	return true
}
