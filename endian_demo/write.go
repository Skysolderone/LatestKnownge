package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// little endian
func main() {
	buf := new(bytes.Buffer)
	var num int32 = 0x12345678
	binary.Write(buf, binary.LittleEndian, num)
	fmt.Printf("little endian:%X\n", buf.Bytes())
	buf.Reset()
	binary.Write(buf, binary.BigEndian, num)
	fmt.Printf("big endian:%X\n", buf.Bytes())
	buf.Reset()
}
