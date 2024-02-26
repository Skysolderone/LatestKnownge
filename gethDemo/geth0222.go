package main

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

// solcjs --abi .\Storage.sol -o build
//
//	abigen --abi Storage.abi --pkg main --type Storage --out Storage.go

const Key = `0x10fbfb7b33FD146460b0249917C203C4ced9dD37`

func main() {
	conn, err := ethclient.Dial("/home/go-ethereum/goerli/geth_ipc")
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(Key), "<stronge password>")
	if err != nil {
		log.Fatal(err)
	}
	address, tx, instance, err := DeployStorage(auth, conn), new(big.Int), "Storage contract in Go!", 0, "Go!")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	// function call on `instance`. Retrieves pending name
	name, err := instance.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)

	store, err := NewStorage(common.HexToAddress("0x21e6fc92f93c8a1bb41e2be64b4e1f88a54d3576"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}
	value, err := store.Retrieve(nil)
if err != nil {
	log.Fatalf("Failed to retrieve value: %v", err)
}
fmt.Println("Value: ", value)
	
}
