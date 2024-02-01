package main

import "fmt"

type Event struct {
	Type string
	Data interface{}
}
type EventStore struct {
	events []Event
}

func (store *EventStore) Append(event Event) {
	store.events = append(store.events, event)
}
func (store *EventStore) GetEvents() []Event {
	return store.events
}

type Account struct {
	id      string
	balance int
	store   *EventStore
}

func NewAccount(id string, store *EventStore) *Account {
	return &Account{
		id:      id,
		balance: 0,
		store:   store,
	}
}

func (account *Account) Deposit(amount int) {
	event := Event{
		Type: "desposot",
		Data: amount,
	}
	account.store.Append(event)
	account.balance += amount
}

func (account *Account) Withdraw(amount int) {
	if account.balance >= amount {
		event := Event{
			Type: "wthdraw",
			Data: amount,
		}
		account.store.Append(event)
		account.balance -= amount
	}

}

func (account *Account) GetBalance() int {
	return account.balance
}

func main() {
	store := &EventStore{}
	account := NewAccount("123", store)
	account.Deposit(100)
	account.Withdraw(50)
	account.Deposit(25)
	events := store.GetEvents()
	for _, event := range events {
		switch event.Type {
		case "deposit":
			amount := event.Data.(int)
			fmt.Printf("Deposited %d\n", amount)
		case "withdraw":
			amount := event.Data.(int)
			fmt.Printf("withdraw %d \n", amount)
		}

	}
	fmt.Printf("Final balance :%d\n", account.GetBalance())
}
