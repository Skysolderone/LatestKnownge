package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	name    string
	age     int
	animals []string
}

func p(s any) { fmt.Printf("%+v\n", s) }
func main() {
	//basic
	x := 10
	xptr := &x
	xuintptr := uintptr(unsafe.Pointer(xptr))
	fmt.Println(*(*int)(unsafe.Pointer(xuintptr)))
	//level1
	var u user
	p(u)
	uNamePtr := (*string)(unsafe.Pointer(&u))
	*uNamePtr = "bradford"
	p(u)
	age := (*int)(unsafe.Add(unsafe.Pointer(&u), unsafe.Offsetof(u.age)))
	*age = 34
	p(u)
	u.animals = []string{"wws", "demo", "k"}
	p(u)
	secondAnimal := (*string)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(u.animals)), unsafe.Sizeof("")))
	p(u)
	*secondAnimal = "calos"
	p(u)

	//level2
	fruits := []string{"apples", "oranges", "bananas", "kansas"}
	start := unsafe.Pointer(unsafe.SliceData(fruits))
	size := unsafe.Sizeof(fruits[0])
	for i := 0; i < len(fruits); i++ {
		p(*(*string)(unsafe.Add(start, uintptr(i)*size)))
	}
	//stringtobyte
	myString := "hello wws"
	byteSlice := unsafe.Slice(unsafe.StringData(myString), len(myString))
	p(byteSlice)
	myBytes := []byte{115, 111, 32, 109, 97, 110, 121, 32, 110,
		101, 97, 116, 32, 98, 121, 116, 101, 115}
	str := unsafe.String(unsafe.SliceData(myBytes), len(myBytes))
	p(str)
}
