package main

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Auth struct {
		AccessSecret string
		AccessExpire int64
		Enable       bool
		Secret       string
	}
}

func main() {
	var c Config
	conf.MustLoad("./config.yaml", &c)
	fmt.Println(c)
}
