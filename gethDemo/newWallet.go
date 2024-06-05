package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// conn, _ := ethclient.Dial("http://127.0.0.1:8545")
	// generate private key
	privatekey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// parse byte
	privatekeybytes := crypto.FromECDSA(privatekey)
	// remove 0x
	fmt.Println("private x :", hexutil.Encode(privatekeybytes)[2:])
	// generate public
	publickey := privatekey.Public()
	publickeyEcdsa, ok := publickey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert publickey")
	}
	publickeybytes := crypto.FromECDSAPub(publickeyEcdsa)
	fmt.Println("public x", hexutil.Encode(publickeybytes)[4:])

	address := crypto.PubkeyToAddress(*publickeyEcdsa).Hex()
	fmt.Println("PUBLIC ADDRESS:", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publickeybytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
