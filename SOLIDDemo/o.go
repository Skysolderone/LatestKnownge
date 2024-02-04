package main

import "fmt"

type PaymentMethod interface {
	Pay()
}

type Payment struct{}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

type CreditCard struct {
	amount float64
}

func (cc CreditCard) Pay() {
	fmt.Printf("Paid %.2f using CreditCard", cc.amount)
}

type PayPal struct {
	amount float64
}

func (pp PayPal) Pay() {
	fmt.Printf("Paid %.2f using PayPal", pp.amount)
}

func main() {
	p := Payment{}
	cc := CreditCard{12.23}
	p.Process(cc)
	pp := PayPal{22.33}
	p.Process(pp)
}
