package main

import "fmt"

type mytest struct {
	name string
}
type test struct {
	wokrder func(map[string]*mytest)
}

var UT = map[string]*test{
	"test": {
		wokrder: Test,
	},
}

func main() {
	test := make(map[string]*mytest, 1)
	ls := mytest{
		name: "test",
	}
	test["test"] = &ls
	UT["test"].wokrder(test)
}

func Test(m map[string]*mytest) {
	fmt.Println(m["test"].name)
}
