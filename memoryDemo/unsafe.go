package main

import (
	"fmt"
	"unsafe"
)

type TestS struct {
	Name string
}

func main() {
	data := TestS{"wws"}

	ls := unsafe.Pointer(&data)
	p := uintptr(ls)
	//
	p2 := unsafe.Pointer(p)
	d2 := (*TestS)(p2)
	fmt.Printf("%#v", d2)
}
