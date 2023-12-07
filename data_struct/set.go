package main

import "fmt"

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}
func (s *set) Add(v string) {
	s.m[v] = struct{}{}
}
func (s *set) Contains(v string) bool {
	_, ok := s.m[v]
	return ok
}
func (s *set) Remove(v string) {
	delete(s.m, v)
}
func main() {
	s := NewSet()
	s.Add("1")
	s.Add("2")
	fmt.Println(s.Contains("1"))
	fmt.Println(s.Contains("2"))
	s.Remove("1")
	fmt.Println(s.Contains("1"))
}
