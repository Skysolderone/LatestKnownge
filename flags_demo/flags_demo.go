package main

import (
	"flag"
	"fmt"
)

func main() {
	wordstr := flag.String("demo", "default", "set demo")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println(*wordstr)
}
