package main

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("my.db", 0o600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
