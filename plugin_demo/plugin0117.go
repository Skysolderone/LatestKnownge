package main

import (
	"fmt"
	"plugin"
)

/*
load plugin
plug, err := plugin.Open("./plugin.dll")

	if err != nil {
		log.Fatal(err)
	}

find and return addr
symplugin, err := plug.Lookup("symbolName")

unload plugin
plug.Close()
*/
type Calculator interface {
	Add(a, b int) int
	Sub(a, b int) int
}
type MyCalculator struct{}

func (m *MyCalculator) Add(a, b int) int {
	return a + b
}
func (m *MyCalculator) Sub(a, b int) int {
	return a - b
}

//go build -buildmode=plugin -o myplugin.so main.go

// func main() {
// fmt.Println("Hello plugin")
// var c Calculator
// if c != nil {
// 	fmt.Println(c.Add(1, 2))
// 	fmt.Println(c.Sub(2, 1))
// }
// }

type Driver interface {
	Name() string
}

func main() {
	p, err := plugin.Open("./test.dll")
	if err != nil {
		panic(err)
	}
	newDriverSymbol, err := p.Lookup("NewDriver")
	if err != nil {
		panic(err)
	}
	newDriverSymbolFunc := newDriverSymbol.(func() Driver)
	newDriver := newDriverSymbolFunc
	fmt.Println(newDriver().Name())
}
