package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

// tls
func main() {
	certificate, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{certificate}}
	listener, err := tls.Listen("tcp", ":1234", config)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello,go")
	})
	http.Serve(listener, nil)

}
