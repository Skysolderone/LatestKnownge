package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

// $ go get -u github.com/tidwall/gjson
const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "name.last")
	fmt.Println(value)
	value2 := gjson.Get(json, "age")
	fmt.Println(value2)
	parseArray()
}

const arrayjson = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

func parseArray() {
	data := gjson.Get(arrayjson, "friends")
	// for _, v := range data.Array() {
	// 	// s:=gjson.Get(json,"")
	// 	fmt.Println(v.String())
	// }
	for _, v := range data.Array() {
		fmt.Println(v.String())
	}
	// fmt.Println(data)
}
