package convert

import "unsafe"

func ByteToSturct(data interface{}) {
}

func StructToByte(data any, n uintptr) []byte {
	ptr := unsafe.Pointer(&data)
	result := unsafe.Slice((*byte)(ptr), n)
	return result
}
