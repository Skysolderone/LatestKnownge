package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	APIKEY := "X9rdbwehnjwb3FNTKbCrmAtKUYu6cnp8lltamaLDkbsQiwmprZagNkSzx7mSQiyVmLoaEPLEjPkJGzKUj87QVg"
	SECRETKEY := "enBARedV8QiRPoYrlgqOmftOF2pyFoBqtXNLYOx9VPYhRkllql81ctRKALjUCaMzKVIXOG2iFAQ6acYtt1WQ"
	APIKEY2 := "X9rdbwehnjwb3FNTKbCrmAtKUYuc6np8IltamaLDkbsQiwmprZagNkSzx7mSQiyVmLoaEPLEjPkJGzKUj87QVg"
	SECRETKEY2 := "enBARedV8QiRPoYrlgqOmft0F2pyFoBqtXNL0x9VPYhRkIql81ctRKALjUCaMzKVIXOG2iFAQ6acYtt1WQ"
	if APIKEY == APIKEY2 {
		fmt.Println("TU")
	}
	if SECRETKEY == SECRETKEY2 {
		fmt.Println("T2")
	}
	ls := md5.New()
	io.WriteString(ls, APIKEY)
	// ls.Write(APIKEY)
	ls2 := md5.New()
	io.WriteString(ls2, APIKEY2)
	fmt.Println(string(ls.Sum(nil)))
	fmt.Println(string(ls2.Sum(nil)))
}
