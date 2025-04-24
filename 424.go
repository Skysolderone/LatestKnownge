package main

type Config struct {
	Auth struct {
		AccessSecret string
		AccessExpire int64
		Enable       bool
		Secret       string
	}
}

func main() {
}
