package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Pointer
	var ptr *int
	isNil := reflect.ValueOf(ptr).IsNil()
	fmt.Println(isNil)
	//interface
	var iface interface{}
	// tye := reflect.TypeOf(&iface)
	// fmt.Println(tye)
	isNil = reflect.ValueOf(&iface).IsNil()

	fmt.Println(isNil)
	//slice
	var s []int
	fmt.Println(reflect.ValueOf(s).IsNil())
	//chan
	var ch chan int
	fmt.Println(reflect.ValueOf(ch).IsNil())
	//map
	var m map[int]string
	fmt.Println(reflect.ValueOf(m).IsNil())
	//func
	var fu func()
	fmt.Println(reflect.ValueOf(fu).IsNil())

}
