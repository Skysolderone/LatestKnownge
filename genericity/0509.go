package main

import "fmt"

func Swap[T any](x, y *T) {
	tmp := *x
	*x = *y
	*y = tmp
}

func main() {
	// s := 1
	// a := 2
	// Swap(&s, &a)
	// fmt.Println(s)
	// fmt.Println(a)
	user := User{ID: 1, Name: "John Doe"}
	product := Product{ID: 2, Name: "Product A", Price: 100.00}

	Print(user)    // 输出：{1 John Doe}
	Print(product) // 输出：{2 Product A 100}
}

type User struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func Print[T any](x T) {
	fmt.Println(x)
}
