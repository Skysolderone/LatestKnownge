package main

import (
	"log"

	"github.com/bytedance/sonic"
)

type AddSymbolReq struct {
	Platform string `json:"platform"`
	Future   uint8  `json:"future"`
	Symbol   string `json:"symbol"`
}

func main() {
	data := []byte(`{"name":"json", "wws":"test1", "wws1":"test2"}`)
	var result map[string]interface{}
	err := sonic.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
	data = []byte{123, 34, 112, 108, 97, 116, 102, 111, 114, 109, 34, 58, 34, 98, 105, 116, 103, 101, 116, 34, 44, 34, 102, 117, 116, 117, 114, 101, 34, 58, 49, 44, 34, 115, 121, 109, 98, 111, 108, 34, 58, 34, 80, 65, 82, 84, 73, 34, 125}
	resut2 := AddSymbolReq{}
	sonic.Unmarshal(data, &resut2)
	log.Println(resut2)
}
