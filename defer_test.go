package test

import (
	"fmt"
	"testing"
)

func test() int {
	i := 0
	defer func() { //2 1
		i++
		fmt.Println("defer 1 :", i)
	}()
	defer func() {
		i++
		fmt.Println("defer 2 :", i)
	}()
	defer func() {
		i++
		fmt.Println("defer 3 :", i)
	}()
	return i
}
func test1() (i int) {
	defer func() {
		i++
		fmt.Println("defer 1 :", i)
	}()
	defer func() {
		i++
		fmt.Println("defer 2 :", i)
	}()
	return i
}

func test2() (i int) {
	i = 5
	defer func() {
		i++
		fmt.Println("defer 1", i)
	}()
	defer func(s int) {
		s++
		fmt.Println("defer 2", s)
	}(i)
	return i
}
func TestDefer(t *testing.T) {
	a := test()
	b := test1()
	c := test2()
	t.Log(a)
	t.Log(b)
	t.Log(c)
}
