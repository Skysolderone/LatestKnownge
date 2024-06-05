package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// exit
		conn.Close()
		ctx.Done()
	}()
	// deploy contract

	// get coin address balance
}
