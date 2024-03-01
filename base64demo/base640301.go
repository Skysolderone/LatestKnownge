package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	var salt string
	_, err := fmt.Scan(&salt)
	if err != nil {
		log.Fatal(err)
	}
	encond := base64.StdEncoding.EncodeToString([]byte(salt))
	fmt.Println(encond)
	pitcurefile, err := os.ReadFile("test.png")
	if err != nil {
		fmt.Println(err)
	}

	reuslt := append([]byte(encond), []byte(pitcurefile)...)
	fmt.Println(reuslt)
	file, err := os.Create("testnewfile")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(pitcurefile)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	// data := []byte("hello world")
	// encond := base64.StdEncoding.EncodeToString(data)
	// fmt.Println(encond)
	// doconde, _ := base64.StdEncoding.DecodeString(encond)
	// fmt.Println(string(doconde))
}
