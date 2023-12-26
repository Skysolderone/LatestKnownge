package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)

	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//reader 即为带缓冲的文件读取器
	//后续的操作都在缓冲中执行
	//readbyte
	char, err := reader.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("read byte is %c", char)
	//readString
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s", line)
	//readbytes
	char1, err := reader.ReadBytes(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read byte is %s", string(char1))

	//按需求设置缓冲大小
	// per, err := bufio.NewReaderSize(file, 1024)
	// if err != nil {
	// 	panic(err)
	// }

}
