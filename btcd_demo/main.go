package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func NewBtcAddress() {
	// generate privatekey
	privatekey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	pubkey := privatekey.PubKey()
	// generate btc address
	addressPkh, err := btcutil.NewAddressPubKey(pubkey.SerializeUncompressed(), &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PRIVATEKEY is :%x", privatekey.Serialize())
	fmt.Printf("Public is :%x", pubkey.SerializeUncompressed())
	fmt.Printf("BTC address is:%x", addressPkh.EncodeAddress())
	// generate btc public key
	pubkeyHash := btcutil.Hash160(pubkey.SerializeCompressed())
	// constructor mutli sign script
	// script, err := txscript.NewScriptBuilder().AddOp(txscript.OP_DUP).
	// 	AddOp(txscript.OP_HASH160).AddData(pubkeyHash).
	// 	AddOp(txscript.OP_EQUALVERIFY).AddOp(txscript.OP_CHECKSIG).Script()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// generate P2SH ADDRESS
	addressP2sh, err := btcutil.NewAddressScriptHashFromHash(pubkeyHash, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("p2sh address is:%s", &addressP2sh.EncodeAddress())
	// generaate bech32 address
	addressbech32sh, err := btcutil.NewAddressScriptHashFromHash(pubkeyHash, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bech32 address is :%s", &addressbech32sh.EncodeAddress())
}

func main() {
	NewBtcAddress()
}
