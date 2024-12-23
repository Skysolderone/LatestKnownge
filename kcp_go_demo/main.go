package main

import (
	"crypto/sha1"
	"io"
	"log"
	"time"

	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
)

func main() {
	// key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	key := pbkdf2.Key([]byte("demo"), []byte("demo2"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	if listern, err := kcp.ListenWithOptions("127.0.0.1:12345", block, 10, 3); err == nil {
		go client()
		for {
			s, err := listern.AcceptKCP()
			if err != nil {
				log.Fatal(err)
			}
			go handlleEcho(s)
		}
	}
}

func client() {
	key := pbkdf2.Key([]byte("demo"), []byte("demo2"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)
	time.Sleep(time.Second)
	conn, err := kcp.DialWithOptions("127.0.0.1:12345", block, 10, 3)
	if err != nil {
		log.Fatal(err)
	}
	for {
		data := time.Now().String()
		buf := make([]byte, len(data))
		log.Println("send", data)
		_, err := conn.Write([]byte(data))
		if err != nil {
			log.Fatal(err)
		}
		if _, err := io.ReadFull(conn, buf); err != nil {
			log.Fatal(err)
		} else {
			log.Println("recv", string(buf))
		}
		time.Sleep(time.Second)
	}
}

func handlleEcho(conn *kcp.UDPSession) {
	buf := make([]byte, 4096)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = conn.Write(buf[:])
		if err != nil {
			log.Println(err)
			return
		}
	}
}
