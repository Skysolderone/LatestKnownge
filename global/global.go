package main

import "github.com/go-resty/resty/v2"

var client = resty.New()

func main() {
	client.Debug = true
	client.SetHeader("str", "ss")
	client.R().Get("https://www.baidu.com/users")
	testtwo()
}

func testtwo() {
	client.SetHeader("hi", "ss")
	client.R().Get("https://www.baidu.com/users")
}
