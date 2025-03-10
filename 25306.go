package main

import "fmt"

type St struct {
	Name string
	Age  int
}

func main() {
	var st *St
	if true {
		st := &St{}
		st.Name = "test"
		fmt.Println(st)
	}
	fmt.Println(st)
}
