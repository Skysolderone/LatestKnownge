package main

import (
	"fmt"
	"reflect"
)

func quickSork1(arr interface{}) {
	value := reflect.ValueOf(arr)
	if value.Kind() != reflect.Slice {
		panic("is not slice")
	}
	qucikSortGenert(arr, 0, value.Len()-1)
}
func qucikSortGenert(arr interface{}, i int, len int) {
	//value := reflect.ValueOf(arr)
	if i < len {
		pivoindex := quckSokc(arr, i, len)
		qucikSortGenert(arr, i, pivoindex-1)
		qucikSortGenert(arr, pivoindex+1, len)
	}
}

func quckSokc(arr interface{}, i, len int) int {
	value := reflect.ValueOf(arr)
	pivot := value.Index(i).Interface()
	left, right := i+1, len
	for left <= right {
		for left <= right && reflect.ValueOf(arr).Index(left).Interface() < pivot {
			left++
		}
		for left <= right && reflect.ValueOf(arr).Index(right).Interface() > pivot { //any 类型无法排序只能比较
			//比较==，!=,排序<,>
			right--
		}
		if left <= right {
			swap(arr, left, right)
			left++
			right--
		}
	}
	swap(arr, i, len)
	return right
}
func swap(arr interface{}, i, len int) {
	value := reflect.ValueOf(arr)
	tmp := value.Index(i).Interface()
	value.Index(i).Set(value.Index(len))
	value.Index(len).Set(reflect.ValueOf(tmp))
}
func main() {
	arr := []int{5, 78, 32, 45, 12, 1}
	quickSork1(arr)
	fmt.Println(arr)
	arr1 := []string{"ba", "ca", "aa", "ga", "fa", "da"}
	quickSork1(arr1)
	fmt.Println(arr1)
}
