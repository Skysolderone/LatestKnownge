package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// r := strings.NewReader("Hello\n")
	// signal
	// lr := io.LimitReader(r, 5)
	// multi
	r := strings.NewReader("first\n")
	r1 := strings.NewReader("second\n")
	r2 := strings.NewReader("three\n")
	r3 := strings.NewReader("four\n")

	lr := io.MultiReader(r, r1, r2, r3)
	// output
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
