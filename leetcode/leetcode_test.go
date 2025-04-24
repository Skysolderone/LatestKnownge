package main

import "testing"

func Test2563(t *testing.T) {
	nums := []int{0, 1, 7, 4, 4, 5}
	l, h := 3, 6
	t.Log(countFairPairs2(nums, l, h))
	nums2 := []int{1, 7, 9, 2, 5}
	l, h = 11, 11
	t.Log(countFairPairs2(nums2, l, h))
}

func Test781(t *testing.T) {
	nums := []int{2, 1, 2, 2, 2, 2, 2, 2, 1, 1}
	t.Log(numRabbits(nums))
	nums2 := []int{10, 10, 10}
	t.Log(numRabbits(nums2))
}
