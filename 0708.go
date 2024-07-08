package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	_ "sync"
)

type t struct {
	Name string
	Age  string
}

func main() {
	//	wg := &sync.WaitGroup{}
	//	wg.Add(10)
	//	for i:=0; i<10; i++ {
	//		go func(count int) {
	//			defer func(){
	//				if r := recover();r != nil {
	//					fmt.Println(r)
	//				}
	//				wg.Done()
	//			}()
	//			panic(count)
	//		}(i)
	//	}
	//	wg.Wait()
	// fmt.Println(time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05"))
	data := []byte("clientId544548clientUserId1timestamp1703128229569")
	key := "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC98QmV1Iji3t4H\ncrtwDGhwwe3e30nljABwHSFvEDyAsRNQI2KRbCbvzQe3DEb2JdoSHNpNdyeSrfPQ\nNh4YifmlLWNc6mUChCWEJgg8UVn2uvCOZgGmYau5T8BDJTq5EI7+lu7xi+zxEJKG\nn+4V1COYrmpi7hssmXqmz51pdA6WlOC2KBH+WJOvvpyypg68kRftM6e51Z/Zw5SI\nzL5nEln4coLnpSo2TH4lqBTzVMGPWFIh7MI4ljwbFgM8zg1JxBFdwgK04D6YSGvt\nSRAUH3EfKoSz8vyX7jpQKsgZcc+4fv6eRJstNtqzyCdY1txYwriCJcq20W6FziPs\nil1dAFz7AgMBAAECggEAKkhHBviKBF7ydBwBR78rXLZ++lcgfx38VvjVIYEuxsMk\ndnl6Dr9Rp54vsnrkLRVeBxwvKIF26LsahGYXqwTD8xY4gU2YcMf0BgflW3AMo9UX\nicuBYIVrkdQEC6509Q+PlXMJ+/FXLI6w35EokelohyuEeUNh/zqnSAz9kKUKPlDB\nYVq8791AokvqKsyvdpLdOub5bEltyQG68TruTb96auqUVcl/Bh6QQVJQsTYmNW4O\n1eBN6Wa/UhPsZV67KAq7Fky4sPvw+f0zp8c7NhAJ3DrWPb2dDlA6GsQN1/7HWE1Q\nCsecGnodnx0qfxO2xyAQ70b0ed6OgCfuHYo1undpSQKBgQDxyt1kKfB56gfXm9ZN\nQuPlrSkvBOSIP8agkZpUG4uLBpjE6Nr2GJ38QXvxRJsSnMb24AYLvxjPRTliOsQ/\nmiTHSTwyWwwCLFyLsvsyO4e4D6/QsgCKBaaKqjcI+ljVY5QJtoTiL+yrfRy5pb1v\nFS94Av6yZ9kboqM9p7f2jzgnHQKBgQDJGjYpMgTg6K0pzI2h3iXJMpQSsuQtDe7G\niTGvt+vPBooAh2HeoaVXTjyiPj+0HSacTppnD1VweIXHOKBQ/2mcAj3aE821hJKM\nR/DHB3mhXd6BQAzTMGUxkeADNGyjYfGr2oG2QB/OOYQcic5y88EnsjNCemZAjij7\nZ8DQdnsg9wKBgQC+BdGiM3ejSHKrVvS/uY8DZOgV5MF59q0yaduguBkdxKxxLBNR\nDx2yLwCqJ2G76exP/tWhi6iC1IN5mHe9CHZZDEw7wuQYHScKPzo9Sx2pYOKqoAne\nbvxow8cGzkDcSyr/yyp53HIif+wAVqoez8aDic8DBKxSlKm3/cSNCBd4XQKBgQCh\nenqi+ZcRQcRKHAZeQ5Lzy7k6iorN4rIomWFcR8MBqa6GXi8oJYQuZpQpcvKUHhQy\nzqsMtjNU7HH/VmjkoTKcIvCne99iMeG64slkB4yTFlzkTBoAOowHEate52mUh8EW\nzY36WeMwJN8AS2eKxlADIAQpDT+t9njR5q3ctkaIYQKBgEK5eNdk41SQfZnAEis2\ndgi/bEx+1bmFHxTIe6iOG4+etyExdOiDDFAsmmPNtY2lq9uKZlEEauQIrg+Ms/lF\nlmnwBDrktJlAZ9poNldD8x61t6AEiIV95ehvZVg0Bi/OJR/99f3qqZkh+S6iCBSF\nWf814KLokaecpa4cGwkUpiyy\n-----END PRIVATE KEY-----"
	res, e := sign(data, []byte(key), crypto.MD5)
	if e != nil {
		fmt.Printf("generate sign false: %s\n", e.Error())
	}
	fmt.Println(res)
}

func sign(srcData []byte, priKey []byte, hash crypto.Hash) (string, error) {
	pkixPrivateKey, err := ParsePKCS8PrivateKey(priKey)
	if err != nil {
		fmt.Println(234)
		return "", err
	}
	h := hash.New()
	_, err = h.Write(srcData)
	if err != nil {
		return "", err
	}
	bytes := h.Sum(nil)
	signedData, err := rsa.SignPKCS1v15(rand.Reader, pkixPrivateKey, hash, bytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signedData), nil
}

func ParsePKCS8PrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil || block.Type != "PRIVATE KEY" { //"PRIVATE KEY" only for PKCS8
		return nil, errors.New("invalid private key")
	}
	privateKeyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPrivateKey, ok := privateKeyInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("invalid private key")
	}
	return rsaPrivateKey, nil
}
