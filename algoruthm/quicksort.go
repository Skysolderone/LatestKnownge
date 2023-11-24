package main

import "fmt"

func quickSork(arr []int) {
	if len(arr) < 1 {
		return
	}
	pivoIndex := partion(arr)
	quickSork(arr[:pivoIndex])
	quickSork(arr[pivoIndex+1:])
}
func partion(arr []int) int {
	pivot := arr[0]
	left, right := 1, len(arr)-1
	for left <= right {
		for left <= right && arr[left] < pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	arr[0], arr[right] = arr[right], arr[0]
	return right
}

func main() {
	arr := []int{4, 7, 89, 52, 1, 34, 24}
	quickSork(arr)
	fmt.Println(arr)
}
