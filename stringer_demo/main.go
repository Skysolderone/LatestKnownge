// main.go
package main

import "fmt"

//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
)

func main() {
	fmt.Println(Placebo)
	fmt.Println(Aspirin)
	fmt.Println(Ibuprofen)
	fmt.Println(Paracetamol)
}
