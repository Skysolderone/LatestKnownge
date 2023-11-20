package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed test.txt test2.txt
var fs embed.FS

// var content []string 单文件

func main() {
	// 单文件
	// fmt.Println(content)
	//多文件
	data, _ := fs.ReadFile("test.txt")
	fmt.Println(string(data))
	data2, _ := fs.ReadFile("test2.txt")
	fmt.Println(string(data2))
}
