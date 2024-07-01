package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"software.sslmate.com/src/go-pkcs12"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{Organization: []string{"wws"}},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	// 自签名证书
	certDer, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(certDer)
	// 证书和私钥打包成pkcs12
	pfsData, err := pkcs12.Encode(rand.Reader, privateKey, template, []*x509.Certificate{template}, "gg123456")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("cert.p12", pfsData, 0o655)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PKCS12 certificate generated successfully")
}
