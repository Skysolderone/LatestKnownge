package main

import (
	"fmt"
	"plugin"
)
//windows 只能使用dll  go build -buildmode=plugin 生成linux的so
func main() {
	p, err := plugin.Open("./plugin.so")
	if err != nil {
		fmt.Println(err)
	}
	greet, err := p.Lookup("Greet")
	if err != nil {
		fmt.Println(err)
	}
	greet.(func(string("wws")))
}
