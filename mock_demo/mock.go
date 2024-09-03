package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	FirstName string `fake:"{firstname}"`
	LastName  string `fake:"{lastname}"`
	Email     string `fake:"{email}"`
	Phone     string `fake:"{phone}"`
	Birthdate string `fake:"{date}"`
}

func main() {
	gofakeit.Seed(0)
	var user User
	gofakeit.Struct(&user)
	fmt.Printf("%#v", user)
	users := []User{}
	gofakeit.Slice(&users)
	fmt.Printf("%#v", users)
}
