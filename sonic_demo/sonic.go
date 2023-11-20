package main

import (
	"log"

	"github.com/bytedance/sonic"
)

func main() {
	data := []byte(`{"name":"json", "wws":"test1", "wws1":"test2"}`)
	var result map[string]interface{}
	err := sonic.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)

	}
	log.Println(result)
}
