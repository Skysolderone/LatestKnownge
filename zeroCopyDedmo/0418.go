package main

import (
	"syscall"
)

func main() {
	n,err:=syscall.Sendfile()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(n)
}
