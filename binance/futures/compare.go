package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("./log.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file1, err := os.Open("./log (2).log")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	res := make(map[string]int)
}
