package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

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

// level1
func encrypt(key []byte, msg string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	handler(err)
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))
	return ciphertext, nil
}

func decrypy(key, ciphertext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	handler(err)
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

func handler(err error) {
	if err != nil {
		panic(err)
	}
}
