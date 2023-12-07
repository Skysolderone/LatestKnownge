package main

import (
	"fmt"
	"log"
)

func quickSork(arr []int) {
	if len(arr) < 1 {
		return
	}
	pivoIndex := partion(arr)
	log.Println(arr[:pivoIndex])
	quickSork(arr[:pivoIndex])
	log.Println(arr[pivoIndex+1:])
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
			log.Println(left, right)//3217
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}

	}
	log.Println(left, right)//1237
	arr[0], arr[right] = arr[right], arr[0]
	log.Println(arr)
	return right
}

func main() {
	arr := []int{3, 2, 7, 1}
	quickSork(arr)
	fmt.Println(arr)
}
