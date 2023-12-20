package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Persion struct {
	Name    string
	Age     int
	Address string
}

func main() {
	//register
	gob.Register(Persion{})
	//eg.
}

func encodeData() {
	file, err := os.Create("encode_data.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	//generate
	person := Persion{
		Name:    "wws",
		Age:     30,
		Address: "string",
	}
	err = encoder.Encode(person)
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	return
}

func decoder() {
	file, err := os.Open("encode_data.gob")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	var deco Persion
	err = decoder.Decode(&deco)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", deco)
}


