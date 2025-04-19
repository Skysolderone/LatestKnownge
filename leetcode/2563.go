package main

import (
	"slices"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	// 效率为O（n2） 数据大性能差
	count := 0
	// slices.Sort(nums)
	for i, v := range nums {
		for j := i + 1; j <= len(nums)-1; j++ {
			if check(v, nums[j], lower, upper) {
				count++
			}
		}
	}
	return int64(count)
}

func check(i, j, l, h int) bool {
	sum := i + j
	if l <= sum && sum <= h {
		return true
	}
	return false
}

func countFairPairs2(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)
	var count int64
	n := len(nums)
	for i := 0; i < n; i++ {
		left := sort.Search(n-1-i, func(j int) bool {
			return nums[i]+nums[j+1+i] >= lower
		})
		riget := sort.Search(n-1-i, func(j int) bool {
			return nums[i]+nums[j+1+i] > upper
		})
		count += int64(riget) - int64(left)
	}
	return count
}
