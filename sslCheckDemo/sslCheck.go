package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

func CheckCertExpiration(domain string) error {
	conn, err := net.DialTimeout("tcp", domain+":443", 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	config := &tls.Config{ServerName: domain}
	tlsConn := tls.Client(conn, config)
	defer tlsConn.Close()
	err = tlsConn.Handshake()
	if err != nil {
		log.Fatal(err)
	}
	certs := tlsConn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Printf("Subject :%s\n", cert.Subject.CommonName)
		fmt.Printf("Expires On:%s\n", cert.NotAfter.Format("2006-01-02"))
		fmt.Printf("Days Left :%d\n", int(cert.NotAfter.Sub(time.Now()).Hours()/24))

	}
	return nil
}

func main() {
	domain := "example"
	err := CheckCertExpiration(domain)
	if err != nil {
		log.Fatal(err)
	}
}
