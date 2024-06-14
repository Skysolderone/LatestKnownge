package main

import "fmt"

func main() {
	open := 500 * 0.80
	makret := 500 * 0.6
	fmt.Println("LOSE:", open-makret)
	open2 := 3000 * 0.6
	sum := open + open2
	fmt.Println("SUM", sum)
	fmt.Println("PRICE:", sum/3500.0)
}
