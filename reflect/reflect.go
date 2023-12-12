package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.TypeOf(a).Kind())
	s := []byte("string")
	b := []byte("int")
	fmt.Println(reflect.Copy(reflect.ValueOf(b), reflect.ValueOf(s))) //复制
	fmt.Println(string(s), string(b))
	if !reflect.DeepEqual(s, b) { //比较
		fmt.Println(true)
	}
	test := []int{1, 2, 3, 4, 5, 6} //交换
	result := reflect.Swapper(test)
	result(4, 5)
	fmt.Println(test)
	//get struct
	Aest := Aest{"wws", 15}
	ls := reflect.TypeOf(Aest)
	fmt.Println(ls.NumField())
	for i := 0; i < ls.NumField(); i++ {
		k := ls.Field(i)
		fmt.Println(k.Index, k.Type, k.Name)
	}
}

type Aest struct {
	Name string
	Age  int
}
