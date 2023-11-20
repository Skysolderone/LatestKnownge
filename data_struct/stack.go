package main

import (
	"fmt"
	"sync"
)

// FILO
//
//	type Node strcut{
//		key int
//		value int
//		front *Node
//		next *Node
//	}
type Stack struct {
	limit int
	head  int
	list  []int
	mute  sync.Mutex
}

func NewStack(limit int) *Stack {
	sta := &Stack{limit: limit}
	sta.list = make([]int, 0)
	return sta
}
func (s *Stack) Push(a int) {
	s.mute.Lock()
	defer s.mute.Unlock()
	if len(s.list) == s.limit {
		return
	}
	s.list = append(s.list, a)
	fmt.Println(s.list)
	s.head = a
}
func (s *Stack) P() {
	fmt.Println(s.list)
}
func (s *Stack) Pop() int {
	s.mute.Lock()
	defer s.mute.Unlock()
	if len(s.list) == 0 {
		return -1
	}
	item := s.list[len(s.list)-1]
	s.list = s.list[0 : len(s.list)-1]
	return item
}

func main() {
	s := NewStack(5)
	s.Push(3)
	s.Push(4)
	s.P()
	// ls := make([]int, 0)
	// ls = append(ls, 4)
	// ls = append(ls, 3)
	// fmt.Println(ls)

}
