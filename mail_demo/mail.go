package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "405620325@qq.com")
	m.SetHeader("To", "wws318310@gmail.com")
	m.SetHeader("Subject", "Hello！")
	m.SetBody("text/plain", "hello wws")
	d := gomail.NewDialer("smtp.qq.com", 587, "405620325@qq.com", "password") //password 应该为邮箱授权码
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}
