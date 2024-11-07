package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// _, err := os.Create("./sre.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// csv.Ge()
	file, err := os.OpenFile("./csv/aaa.csv", os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	result, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	writer := csv.NewWriter(file)
	err = writer.Write([]string{"sss"})
	if err != nil {
		fmt.Println(err)
	}
	writer.Flush()
}
