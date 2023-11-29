package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[2:4]
	change(s1)
	fmt.Println(s)
	fmt.Println(s1)

}

func change(s []int) {
	s = append(s, 3)
	fmt.Println(s)
}
