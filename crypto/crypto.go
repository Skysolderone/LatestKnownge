package main

import "fmt"

func main() {
	//basic
	//对称加密
	key := generateSymmetricKey()
	encryptedText := encryptSymmetric("Hello,go", key)
	decryptedText := decryptedText(encryptedText, key)
	//非对称加密
	publicKey, privateKey := generateKeyPair()
	encryptedMessage := encryptAsymmetric("Hello,go", publicKey)
	decryptedMessage := decryptAsymetric(encryptedMessage, privateKey)
	fmt.Println("Sym :", decryptedText)
	fmt.Println("Asym", decryptedMessage)
}
