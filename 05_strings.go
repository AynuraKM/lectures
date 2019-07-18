package main

import "fmt"

type Strings struct {
	str []string
}

func (s *Strings) Len() int {
	return len(s.str)
}

func (s *Strings) Swap(i, j int) []string {
	if s.Less(i, j) {
		s.str[i], s.str[j] = s.str[j], s.str[i]
		return s.str
	}
	return s.str
}

func (s *Strings) Less(i, j int) bool {
	if s.str[i] > s.str[j]{
		return false
	}
	return true
}

func main() {
	s := Strings{[]string{"Rusya", "Kolya", "Summ", "left", "oreng"}}
	fmt.Println(s.Less(1, 4))
	fmt.Println(s.Len())
	fmt.Println(s.Swap)

}
