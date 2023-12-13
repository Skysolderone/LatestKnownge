package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const INT_SIZE = int(unsafe.Sizeof(0)) //64bit 8byte
// judge local compater binary
func SystemEdian() {
	var i = 0x01020304
	fmt.Println("&i:", &i)
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0x04 {
		fmt.Println("little")
	} else {
		fmt.Println("big")
	}
	fmt.Printf("0x%x,%v\n", bs[0], &bs[0])
	fmt.Printf("0x%x,%v\n", bs[1], &bs[1])
	fmt.Printf("0x%x,%v\n", bs[2], &bs[2])
	fmt.Printf("0x%x,%v\n", bs[3], &bs[3])
}

func testBigEdian() {
	var testint int32 = 0x01020304
	fmt.Printf("%d use big edian:\n", testint)
	testBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testint))
	fmt.Println(testBytes)
	fmt.Printf("%x\n", testBytes)
	convInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("convert int32 :%d\n", convInt)

}
func testLittleEdian() {
	var testint int32 = 0x01020304
	fmt.Printf("%d use little edian:\n", testint)
	testBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testint))
	fmt.Println(testBytes)
	fmt.Printf("%x\n", testBytes)
	convInt := binary.LittleEndian.Uint32(testBytes)
	fmt.Printf("convert int32 :%d\n", convInt)

}

// 字节序大端小端
func main() {
	SystemEdian()
	testBigEdian()
	testLittleEdian()
}
