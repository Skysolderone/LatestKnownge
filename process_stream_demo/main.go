package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// stream data write into http request body
func main() {
	// mock data
	largeData := strings.Repeat("Hello world", 1000000)
	// convert io.Reader
	render := strings.NewReader(largeData)
	// create http request,set reader body
	req, err := http.NewRequest("POST", "http://example.com/upload", render)
	if err != nil {
		log.Fatal(err)
	}
	// set content-type header
	req.Header.Set("Content-Type", "text/plain")
	// send req
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status %#v", resp.Status)
	fmt.Printf("resp body %#v", string(body))
}
