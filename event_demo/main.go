package main

import (
	"fmt"
	"log"
	"net"
	"v1/event_demo/events"
)

func handleConnect(conn net.Conn, dispacher *events.Dispatcher) {
	dispacher.DisPatch(events.NetEvent{Type: "Connect", Conn: conn})
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			dispacher.DisPatch(events.NetEvent{Type: "disConnect", Conn: conn})
			return
		}
		dispacher.DisPatch(events.NetEvent{Type: "receive", Conn: conn, Message: buffer[:n]})
	}
}
func startTcpServer(addr string, dispacher *events.Dispatcher) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnect(conn, dispacher)

	}

}
func main() {
	dispacher := events.NewDispatcher()
	//register
	// listener := events.NewEventListener(func(e events.NetEvent) {
	// 	fmt.Printf("Event received:%s\n", e.Type)
	// })
	// dispacher.RegisterListener("Connect", listener)
	// dispacher.RegisterListener("disConnect", listener)
	// dispacher.RegisterListener("receive", listener)
	dispacher.RegisterListener("Connect", func(e events.NetEvent) {
		//netEvent := e.(events.NetEvent)
		fmt.Println("New connection:", e.Conn.RemoteAddr())
	})

	dispacher.RegisterListener("receive", func(e events.NetEvent) {
		//netEvent := e.(events.NetEvent)
		fmt.Printf("Received data: %s\n", string(e.Message))
	})

	dispacher.RegisterListener("disConnect", func(e events.NetEvent) {
		//netEvent := e.(events.NetEvent)
		fmt.Println("Connection closed:", e.Conn.RemoteAddr())
	})
	//dispacher.DisPatch(events.NetEvent{Type: "tcpserver"})
	startTcpServer(":8080", dispacher)
}
