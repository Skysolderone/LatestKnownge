package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"sync"

	"golang.org/x/time/rate"
)

type Liq struct {
	ATR        float64
	Resistance float64
	Support    float64
}

var (
	okexMap     = make(map[string]Liq, 0)
	okexRw      sync.RWMutex
	OkxLimitMap = make(map[string]*rate.Limiter, 0)
)

func getLiq(exchange string, sym string) Liq {
	limit, ok := OkxLimitMap[sym]
	if !ok {
		okexRw.RLock()
		OkxLimitMap[sym] = rate.NewLimiter(1/1.0, 1)
		okexRw.RUnlock()
	}

	if v, ok := okexMap[sym]; ok && !limit.Allow() {
		log.Println("------limit ")
		fmt.Println(sym)
		return v
	} else {
		log.Println("------init ")
		fmt.Println(sym)
		s := Liq{ATR: rand.Float64(), Resistance: rand.Float64(), Support: rand.Float64()}
		okexRw.RLock()
		okexMap[sym] = s
		okexRw.RUnlock()
		return s
	}
}

var s = map[uint8]string{
	123: "123",
}

func main() {
	// for {
	// 	for i := range 3 {
	// 		s := fmt.Sprintf("%d-usdt", i)
	// 		res := getLiq("okex", s)

	// 		fmt.Println(res)
	// 		time.Sleep(time.Second * 1)
	// 	}
	// }
	key := s[5]

	if key == "" {
		key = "nil"
	}
	fmt.Println(key)
}
