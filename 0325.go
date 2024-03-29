package main

import (
	"fmt"
	"strings"
)

func main() {
	// res, err := http.Get("https://t.me/+S1sKwgubYGdjNjQ5")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// // if len(res.Header["X-Frame-Options"]) != 0 && res.Header["X-Frame-Options"][0] == "ALLOW-FROM https://web.telegram.org" {
	// // 	log.Println("TRUE")
	// // }
	// log.Println(res)
	// // log.Println(res.Header["X-Frame-Options"][0])
	s := "tusdtT"
	if strings.HasSuffix(s, "T") {
		s = strings.TrimRight(s, "T")
		fmt.Println(s)
	}
}
