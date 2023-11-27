package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	//brocaset
	brocast := NewbBro()
	go brocast.Run()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		go handle(brocast, conn)

	}

}

// brocast
type Brocast struct {
	clients  map[chan<- string]struct{} //all write chan
	joins    chan chan<- string         //new connect chan
	leaves   chan chan<- string         //disconnect chan
	messages chan string
}

func NewbBro() *Brocast {
	return &Brocast{
		clients:  make(map[chan<- string]struct{}),
		joins:    make(chan chan<- string),
		leaves:   make(chan chan<- string),
		messages: make(chan string),
	}
}
func (b *Brocast) Run() {
	for {
		select {
		case joins := <-b.joins:
			b.clients[joins] = struct{}{}
		case leave := <-b.leaves:
			delete(b.clients, leave)
			close(leave)
		case msg := <-b.messages:
			for client := range b.clients {
				client <- msg
			}
		}
	}
}

// client
type ClientType struct {
	conn    net.Conn
	message chan string
}

func NewClient(conn net.Conn) *ClientType {
	return &ClientType{
		conn:    conn,
		message: make(chan string),
	}
}
func (ct *ClientType) Write(msg string) {
	ct.message <- msg
}

// func (ct *ClientType)Read(){}
func (ct *ClientType) Run() {
	for msg := range ct.message {
		_, err := fmt.Fprintf(ct.conn, msg)
		if err != nil {
			panic(err)
		}

	}
}

func handle(bro *Brocast, conn net.Conn) {
	client := NewClient(conn)
	go client.Run()
	join := make(chan string)
	bro.joins <- join
	defer func() {
		close(join)
		bro.leaves <- join
	}()
	client.Write("welcome")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		bro.messages <- message
	}
	fmt.Println("process finish:", conn.RemoteAddr())
}
