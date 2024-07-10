package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	littleEndianData := []byte{0x78, 0x56, 0x34, 0x12}
	bigEndianData := []byte{0x12, 0x34, 0x56, 0x78}
	var LittleNum, BigNUM int32
	binary.Read(bytes.NewReader(littleEndianData), binary.LittleEndian, &LittleNum)
	fmt.Printf("%X\n", LittleNum)
	binary.Read(bytes.NewBuffer(bigEndianData), binary.BigEndian, &BigNUM)
	fmt.Printf("%X\n", BigNUM)
}
