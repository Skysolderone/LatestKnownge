package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	var requestGroup singleflight.Group
	// localhost:8080/normal
	http.HandleFunc("/normal", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		status, err := githubStatus()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("/github handler request: processing time :%+v|stacus :%q", time.Since(start), status)
		fmt.Fprintf(w, "GITHUB status :%q", status)
	})
	http.HandleFunc("/singleflight", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
			return githubStatus()
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		status := v.(string)
		log.Printf("/github handler request: processing time :%+v|stacus :%q|shared %t", time.Since(start), status, shared)
		fmt.Fprintf(w, "GITHUB status :%q", status)
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func githubStatus() (string, error) {
	time.Sleep(1 * time.Second)
	log.Println("call githubstatus")
	resp, err := http.Get("https://api.github.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github response: %s", resp.Status)
	}
	return resp.Status, err
}
