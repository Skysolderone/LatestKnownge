package main

import (
	"fmt"

	"v1/structToMutx/p"
)

func main() {
	s := p.NewMutex()
	a := 1
	go func(int) {
		s.Lock()
		a = 2
		s.UnLock()
		fmt.Println(a)
	}(a)
	go func(int) {
		s.Lock()
		a = 3
		s.UnLock()
		fmt.Println(a)
	}(a)
	fmt.Println(a)
	for {
	}
}
