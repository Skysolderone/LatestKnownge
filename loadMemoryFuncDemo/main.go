package main

import (
	"fmt"
	"reflect"
)

func main() {
	funcvalue := reflect.ValueOf(482040)
	// a := 1
	// b := 2
	result := funcvalue.Call(nil)
	fmt.Println(result)
}
