package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ifs interface{}
	fmt.Println(reflect.ValueOf(ifs).IsValid())

}
