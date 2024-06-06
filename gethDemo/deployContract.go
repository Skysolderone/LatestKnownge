package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"v1/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("dial", err)
	}
	defer conn.Close()
	account := common.HexToAddress("0x3D08fa401f2B52cd489e6201F5b1aBC517b678F7")
	balance, _ := conn.BalanceAt(ctx, account, nil)
	fmt.Println(balance)

	//  remove 0x
	// privatekey
	privatekey, err := crypto.HexToECDSA("9d5d5073e60bce2db35fb38c433959669a0b4542bcf6ec6fb8232cf5cadb35c2")
	if err != nil {
		log.Fatal("get privateKey:", err)
	}
	fmt.Println("PRIVATE KEY:", privatekey)
	publicKey := privatekey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error parseing public:")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("PUBLIC ADDRESS:", fromAddress)
	balance2, _ := conn.BalanceAt(ctx, fromAddress, nil)
	fmt.Println("BA", balance2)
	nonce, err := conn.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal("nonce", err)
	}
	gasPrice, err := conn.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal("gas", err)
	}

	// deploy
	auth := bind.NewKeyedTransactor(privatekey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	// input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	_ = instance
}
