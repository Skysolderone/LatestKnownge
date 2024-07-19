package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws/{channel}", handleWebsocket)
	fs := http.FileServer(http.Dir("./web"))
	r.PathPrefix("/").Handler(fs)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channel := vars["channel"]
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	sub := rdb.Subscribe(ctx, channel)
	defer sub.Close()
	ch := sub.Channel()
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			if err := rdb.Publish(ctx, channel, string(msg)).Err(); err != nil {
				log.Println("Publis err:", err)
				return
			}
		}
	}()
	for msg := range ch {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			log.Println("write err:", err)
			return
		}
	}
}
