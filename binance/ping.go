package main

import "net/http"

func main() {
	resp, err := http.Get("https://api.binance.us/api/v3/ping")
	if err != nil {
		panic(nil)
	}
	defer resp.Body.Close()
}
