package main

import (
	"reflect"
	"unsafe"
)

// none，不显示任何 goroutine 堆栈信息
// single，默认级别，显示当前 goroutine 堆栈信息
// all，显示所有 user （不包括 runtime）创建的 goroutine 堆栈信息
// system，显示所有 user + runtime 创建的 goroutine 堆栈信息
// crash，和 system 打印一致，但会生成 core dump 文件（Unix 系统上，崩溃会引发 SIGABRT 以触发core dump）
// 如果我们将 GOTRACEBACK 设置为 system
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Modify() {
	a := "hello"
	b := String2Bytes(a)
	b[0] = 'H'
}

func main() {
	Modify()
}
