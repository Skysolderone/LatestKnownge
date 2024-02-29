package main

import (
	"encoding/json"
	"fmt"
	"log"
)
//gemini
func main() {
	jsonStr := `{
        "name": "John Doe",
        "age": 30,
        "address": {
            "street": "123 Main Street",
            "city": "New York",
            "state": "NY",
            "zip": "10001"
        }
    }`

	// 使用 interface{}

	// var data interface{}
	// err := json.Unmarshal([]byte(jsonStr), &data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 遍历 data
	// switch v := data.(type) {
	// case map[string]interface{}:
	// 	for k, v := range v {
	// 		fmt.Println(k, v)
	// 	}
	// case []interface{}:
	// 	for _, v := range v {
	// 		fmt.Println(v)
	// 	}
	// }

	// 使用 json.RawMessage

	var data json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		log.Fatal(err)
	}

	// 访问 data
	// fmt.Println(data)

	// 将 data 解析为特定的类型
	var v interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 v
	switch v := v.(type) {
	case map[string]interface{}:
		for _, v := range v {
			// 遍历 address
			for k, v := range v.(map[string]interface{}) {
				fmt.Println(k, v)
			}
		}
	case []interface{}:
		for _, v := range v {
			fmt.Println(v)
		}
	}
}
