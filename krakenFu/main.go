package main

import (
	"fmt"
	"sort"
)

type ponuse struct {
	Id   uint
	Name string
}

func main() {
	s := ponuse{Id: 1, Name: "w"}
	s1 := ponuse{Id: 2, Name: "s"}
	s2 := ponuse{Id: 3, Name: "e"}
	lis := make([]ponuse, 0)
	lis = append(lis, s2, s1, s)
	fmt.Println(lis)
	sort.Slice(lis, func(i, j int) bool {
		return lis[i].Id < lis[j].Id
	})
	fmt.Println(lis)
}
