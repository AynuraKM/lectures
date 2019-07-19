package main
import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	s := Strings{"a", "ab", "abc", "abcd", "abcd"}
	fmt.Println(s.Len())
}

func TestStrings_Swap(t *testing.T) {
	s := Strings{"a", "ab", "abc", "abcd", "abcd"}
	s.Swap(0, 1)
	fmt.Println(s)
}
func TestStrings_Less(t *testing.T) {
	s := Strings{"a", "ab", "abc", "abcd", "abcd"}
	fmt.Println(s.Less(0, 1))
}
