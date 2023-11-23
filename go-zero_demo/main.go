package main

import "log"

func main() {
	a := []int{1, 7, 9}
	b := []int{2, 3, 8}
	c := make([]int, 0)
	n := 0
	max := make([]int, 0)
	if len(a) > len(b) {
		n = len(b)
		max = a
	} else {
		n = len(a)
		max = b
	}

	for i := 0; i < n; i++ {
		if b[i] > a[i] {
			c = append(c, a[i])
			c = append(c, b[i])
		} else {
			c = append(c, b[i])
			c = append(c, a[i])
		}
	}
	if len(a) != len(b) {
		c = append(c, max[len(max)-1])
	}
	log.Println(c)

}
