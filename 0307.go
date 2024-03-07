package main

import (
	"fmt"
)

type User struct {
	Name  string
	Role  string
	Id    int
	Count int
}

func main() {
	users := []User{{Name: "modi", Role: "admin", Id: 7, Count: 5}, {Name: "zhangqiang", Role: "admin", Id: 8, Count: 5}, {Name: "songyuan", Role: "admin", Id: 9, Count: 5}}
	for i := range users {
		obj := users[i]
		fmt.Println(obj)
	}
}
