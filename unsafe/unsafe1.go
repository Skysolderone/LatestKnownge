package main

import (
	"fmt"
	"unsafe"
)

type People struct {
	Name string
	Age  int
}

func main() {
	p := People{
		"wws", 15,
	}
	i := unsafe.Sizeof(p)
	j := unsafe.Alignof(p)
	k := unsafe.Offsetof(p.Age)
	fmt.Println("byte sum :", i)
	fmt.Println("align:", j)
	fmt.Println("offset:", k)
	fmt.Printf("address %p:", &p)

	h := unsafe.Pointer(&p)
	fmt.Println(h)
	//
	var x struct {
		a bool
		b int16
		c []int
	}

	/** unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞. */

	/** uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b) 指针的运算 */
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
	//不能分开写
	// NOTE: subtly incorrect!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb = (*int16)(unsafe.Pointer(tmp))
	*pb = 42
}
