package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type obj struct {
	Symbols []struct{}
}

func main() {
	body, err := os.ReadFile("./config.json")
	if err != nil {
		panic("parse err")
	}
	// defer body.Close()
	// obj := io.ReadAll()
	// fmt.Println(string(body))
	// ls := make(map[string]interface{}, 0)
	ls := obj{}
	err = json.Unmarshal(body, &ls)

	if err != nil {
		panic("parse err")
	}
	fmt.Println(ls)
	// if ls.([]inference) {
	// 	objs := ls["symbols"]
	// }
	// for _, v := range ls["symbols"] {
	// 	fmt.Println(v.Symbol[:len(v.Symbol)-4])
	// }
}
