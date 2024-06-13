package main

import (
	"context"
	"fmt"
	"log"

	"v1/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	conn, err := ethclient.Dial("https://data-seed-prebsc-1-s1.bnbchain.org:8545")
	if err != nil {
		log.Fatal("dial", err)
	}
	defer conn.Close()
	// account := common.HexToAddress("0x8b8A0EedC77e05D03f224F1e7218F5Fa10B2922B")
	// balance, _ := conn.BalanceAt(ctx, account, nil)
	// fmt.Println(balance)

	//  remove 0x
	// privatekey
	privatekey, err := crypto.HexToECDSA("metamask privatekey")
	if err != nil {
		log.Fatal("get privateKey:", err)
	}
	// fmt.Println("PRIVATE KEY:", privatekey)
	// publicKey := privatekey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error parseing public:")
	// }
	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// fmt.Println("PUBLIC ADDRESS:", fromAddress)
	// balance2, _ := conn.BalanceAt(ctx, fromAddress, nil)
	// fmt.Println("BA", balance2)
	// nonce, err := conn.PendingNonceAt(ctx, fromAddress)
	// if err != nil {
	// 	log.Fatal("nonce", err)
	// }
	// gasPrice, err := conn.SuggestGasPrice(ctx)
	// if err != nil {
	// 	log.Fatal("gas", err)
	// }
	chainid, _ := conn.ChainID(ctx)
	// deploy
	auth, err := bind.NewKeyedTransactorWithChainID(privatekey, chainid)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// input := "1.0"
	address, tx, instance, err := contract.DeployContract(auth, conn)
	if err != nil {
		log.Fatal(err)
	}
	chatge, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(chatge)
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	_ = instance
}
