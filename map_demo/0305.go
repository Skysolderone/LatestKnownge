package main

import "log"

var once = make(map[int]uint8, 0)

func main() {
	if res, ok := once[5]; !ok {
		log.Println("not exist")
		once[5] = 4
	} else {
		log.Println(res)
	}

	if res, ok := once[5]; !ok {
		log.Println("not exist")
	} else {
		log.Println(res)
	}
}
