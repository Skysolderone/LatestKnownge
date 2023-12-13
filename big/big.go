package main

import (
	"fmt"
	"math/big"
)

func main() {
	//int
	num := big.NewInt(1234567890)
	fmt.Println(num)
	//rat 分数 有理数
	num1 := big.NewRat(1, 2)
	num2 := big.NewRat(3, 4)
	sum := new(big.Rat)
	sum.Add(num1, num2)
	fmt.Println(sum)//5/4

}
