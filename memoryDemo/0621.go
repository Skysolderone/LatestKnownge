package main

import (
	"fmt"
	"unsafe"
)

type data1 struct {
	A int
	B bool
	C string
}

func main() {
	slice := []byte{1, 2, 3, 4}
	slice = append(slice, 5)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	s := data1{
		A: 11,
		B: true,
		C: "wws",
	}

	result := unsafe.Slice((*byte)(unsafe.Pointer(&s)), unsafe.Sizeof(s))
	// bytdat := convert.StructToByte(s, n)

	var newdata data1
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&newdata)), unsafe.Sizeof(s)), result)
	fmt.Printf("%#v", newdata)
}
