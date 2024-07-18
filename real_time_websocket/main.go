package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v4"
)

var (
	connPool *pgx.Conn

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func initDB() {
	var err error
	connPool, err = pgx.Connect(context.Background(), "postgres://postgres:gg123456@172.22.0.1:5432/websocket")
	if err != nil {
		log.Println(err)
	}
	log.Println("CONNECT postgresDb success")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}()
	log.Println("Listerning for database notifications...")
	if _, err := connPool.Exec(context.Background(), "LISTEN events"); err != nil {
		log.Fatalf("fail")
	}
	for {
		notifaction, err := connPool.WaitForNotification(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Received notification: %v", notifaction.Payload)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(notifaction.Payload)); err != nil {
			log.Println("WriteMessage error:", err)
			return
		}
	}
}

func main() {
	initDB()
	defer connPool.Close(context.TODO())
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
