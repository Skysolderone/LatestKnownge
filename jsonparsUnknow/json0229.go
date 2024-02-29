package main

import (
	"encoding/json"
	"fmt"
)

func parseJSON(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	// 假设你有一个 JSON 字符串
	jsonString := `{"name": "John", "age": 30, "city": {"name": "New York", "population": 8000000}}`

	// 将 JSON 字符串转换为字节数组
	jsonData := []byte(jsonString)

	// 解析 JSON 数据
	parsedData, err := parseJSON(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 打印解析结果
	fmt.Printf("%+v\n", parsedData)

	// 访问解析结果中的数据
	if name, ok := parsedData["name"].(string); ok {
		fmt.Println("Name:", name)
	}

	if city, ok := parsedData["city"].(map[string]interface{}); ok {
		if cityName, ok := city["name"].(string); ok {
			fmt.Println("City Name:", cityName)
		}
		if pos, ok := city["population"].(float64); ok {
			fmt.Println("population Name:", pos)
		}
	}
}
