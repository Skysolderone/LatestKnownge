package test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/test-go/testify/require"
)

func Test_0RTT连接建立(t *testing.T) {
	// root 私钥，服务器保存
	root, err := ecdh.P256().GenerateKey(rand.Reader)
	require.NoError(t, err)

	type Data struct {
		Pub []byte // client ecdh pub key
		Sec []byte // encrypted secret and valid pub key
	}

	// 需求：交换secret并验证用户
	var p Data // 表示网络传输的数据

	{ // client
		builtin := root.PublicKey()     // root公钥, 程序内置
		userid := []byte("hello world") // 要求不暴露，或者可以认为其是username&pwd拼接成的，通常在程序登录时获取

		pri, err := ecdh.P256().GenerateKey(rand.Reader)
		require.NoError(t, err)

		pub := pri.PublicKey().Bytes()

		secret, err := pri.ECDH(builtin)
		require.NoError(t, err)
		block, err := aes.NewCipher(secret)
		require.NoError(t, err)
		gcm, err := cipher.NewGCM(block)
		require.NoError(t, err)

		p = Data{
			Sec: gcm.Seal(nil, pub[:12], userid, pub),
			Pub: pub,
		}

		fmt.Println("clinet secret", secret)
	}

	fmt.Println()
	fmt.Println("data size", len(p.Pub)+len(p.Sec))
	fmt.Println()

	{ // server

		clientPub, err := ecdh.P256().NewPublicKey(p.Pub)
		require.NoError(t, err)

		secret, err := root.ECDH(clientPub)
		require.NoError(t, err)

		block, err := aes.NewCipher(secret)
		require.NoError(t, err)
		gcm, err := cipher.NewGCM(block)
		require.NoError(t, err)

		userid, err := gcm.Open(nil, p.Pub[:12], p.Sec, p.Pub)
		require.NoError(t, err)

		fmt.Println("client secret", secret)
		fmt.Println("user id:", string(userid)) // 验证user id
	}
}
