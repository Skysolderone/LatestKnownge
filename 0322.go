package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://t.me/+S1sKwgubYGdjNjQ5")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if len(res.Header["X-Frame-Options"]) != 0 && res.Header["X-Frame-Options"][0] == "ALLOW-FROM https://web.telegram.org" {
		log.Println("TRUE")
	}
	log.Println(res.Body)
	// log.Println(res.Header["X-Frame-Options"][0])
}
