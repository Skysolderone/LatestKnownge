package main

import "fmt"

type User struct {
	Age int
	Id  int
}

func main() {
	// ls := []User{{Age: 5, Id: 0}, {Age: 6, Id: 0}}
	// sort.Slice(ls, func(i, j int) bool {
	// 	return ls[i].Age < ls[j].Age
	// })
	// for i := range ls {
	// 	ls[i].Id = i + 1
	// 	// fmt.Println(v.Id)
	// }
	// fmt.Println(ls)
	ls := make(map[string]User, 0)
	u := User{}
	u.Age = 5
	u.Id = 4
	ls["test"] = u
	if o, ok := ls["test"]; ok {
		fmt.Println(o)
		o.Age = 3
		
	}
}
