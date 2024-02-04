package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) MakeSound() {
	fmt.Println("Animal sound")
}

type Bird struct {
	Animal
}

func (b Bird) MakeSound() {
	fmt.Println("Chirp chirp")
}

type AnimalBehavior interface {
	MakeSound()
}

func MakeSound(ab AnimalBehavior) {
	ab.MakeSound()
}
func main() {
	a := Animal{}
	b := Bird{}
	MakeSound(a)
	MakeSound(b)
}
