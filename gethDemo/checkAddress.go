package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is val
	// need network
	conn, _ := ethclient.Dial("http://127.0.0.1:8545")
	ctx := context.Background()
	// check address is account or contract
	address := common.HexToAddress("0x35A42428a5446E35158b90D6339F8eAaEf95c272")
	bytecode, err := conn.CodeAt(ctx, address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	// account(false) contract(true)
	fmt.Printf("IS contract:%v\n", isContract)
}
