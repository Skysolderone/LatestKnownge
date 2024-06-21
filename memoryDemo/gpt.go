package main

import (
	"fmt"
	"unsafe"
)

type nested struct {
	X float64
	Y float64
}

type data struct {
	A int
	B bool
	C string
	D nested
}

func main() {
	// 创建一个data类型的实例
	s := &data{
		A: 11,
		B: true,
		C: "Hello, World!",
		D: nested{
			X: 3.14,
			Y: 2.71,
		},
	}

	// 打印原始结构体
	fmt.Printf("Original Struct: %+v\n", *s)

	// 获取data实例的大小
	size := unsafe.Sizeof(*s)
	// 将指针转换为unsafe.Pointer
	ptr := unsafe.Pointer(s)
	// 创建一个byte切片来访问data实例的字节表示
	byteSlice := unsafe.Slice((*byte)(ptr), size)

	// 打印字节切片
	fmt.Println("Byte Slice:", byteSlice)

	// 创建一个新的data实例，用于反转字节切片
	var newData data
	// 获取newData的指针
	newPtr := unsafe.Pointer(&newData)
	// 将字节切片复制到新的结构体实例中
	copy(unsafe.Slice((*byte)(newPtr), size), byteSlice)

	// 单独处理字符串字段
	// newData.C = s.C

	// 打印反转后的结构体
	fmt.Printf("Reversed Struct: %+v\n", newData)
}
