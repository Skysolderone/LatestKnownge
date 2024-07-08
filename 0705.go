package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	// callbackUrl := "https://api.utrading.io/oauth/callback/bitget"
	// params := url.Values{}
	// params.Set("redirectUrl", callbackUrl)
	// fmt.Println(params.Encode())
	key, _ := readPrivateKey("bitget_key.pem")
	sign, err := key.Sign(rand.Reader, []byte("wws"), crypto.MD5)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(sign)
}

func readPrivateKey(filename string) (*rsa.PrivateKey, error) {
	// 读取私钥文件内容
	privateKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解码 PEM 格式的私钥块
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil {
		return nil, errors.New("error decoding PEM private key block")
	}

	// 解析 DER 编码的私钥块为 *ecdsa.PrivateKey
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}

	// 使用 privateKey 进行操作，比如签名等
	return privateKey, nil
}
