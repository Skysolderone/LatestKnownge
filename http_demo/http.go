package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, h *http.Request) {
		w.Write([]byte("hello "))
	})
	http.ListenAndServe(":1234", nil)
}
