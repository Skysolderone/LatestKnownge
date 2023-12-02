package main

import (
	"fmt"
	"reflect"
)

type SampleInterface interface {
	Print()
}

type SampleStruct struct{}

func (s *SampleStruct) Print() {
	fmt.Println("sample")
}

func main() {
	//type
	//var test int
	//基类型匹配
	var test1 int64 = 42
	// type1 := reflect.TypeOf(test)
	type1 := reflect.TypeOf(test1)
	if type1.Kind() == reflect.Int {
		fmt.Println("true") //int64 不是int
	}
	//接口类型匹配
	sample := SampleStruct{}
	sampleType := reflect.TypeOf(sample)
	if reflect.TypeOf((*SampleInterface)(nil)).Elem().AssignableTo(sampleType) {
		fmt.Println("true")
	}
}
