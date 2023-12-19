package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CustomData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// pares file
func ParseDataFile(filename string) (*CustomData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var customData CustomData
	err = json.Unmarshal(data, &customData)
	if err != nil {
		return nil, err
	}
	return &customData, nil
}

// serialize data to file
func SerializeDatatoFile(data *CustomData, filename string) error {
	jsondata, err := json.MarshalIndent(data, "", "     ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, jsondata, 0644)
	if err != nil {
		return err
	}
	return nil
}

// basic
func main() {
	data, err := ParseDataFile("data.json")
	if err != nil {
		fmt.Println("error parse data:", err)
		return
	}
	fmt.Printf("parse data:%+v\n", data)
	//modtify data
	data.Value = 42
	//save data to file
	err = SerializeDatatoFile(data, "modtify_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data save to file")

}
