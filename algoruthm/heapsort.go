package main

import "fmt"

func heapsort(s []int) {
	length := len(s) - 1
	for k := length / 2; k >= 1; k-- {
		sift(s, k, length)
	}
	for length > 1 {
		swap(s, 1, length)
		length--
		sift(s, 1, length)
	}
}
func sift(s []int, k, length int) {
	for {
		i := 2 * k
		if i > length {
			break
		}
		if i < length && s[i+1] > s[i] {
			i++
		}
		if s[k] == s[i] {
			break
		}
		swap(s, k, i)
		k = i
	}
}
func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	s := []int{-1, 47, 8, 6, 1, 2, 3}
	fmt.Println(s[1:])
	heapsort(s)
	fmt.Println(s[1:])

}
