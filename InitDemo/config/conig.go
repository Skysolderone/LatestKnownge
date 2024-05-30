package config

import "fmt"

func init() {
	Test = "string"
	fmt.Println(Test)
}

var Test = "test"
