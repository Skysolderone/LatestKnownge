package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bytedance/sonic"
	gate "github.com/gateio/gatews/go"
)

func main() {
	ws, err := gate.NewWsService(nil, nil, gate.NewConnConfFromOption(&gate.ConfOptions{
		Key:           "e45bccfa36e318360d6012027a3741c6",
		Secret:        "2bffdfd9173fe718f8cb7ba6ba67ca99967a1d3acf0936f3004ea3382ce49f41",
		MaxRetryConn:  10, // default value is math.MaxInt64, set it when needs
		SkipTlsVerify: false,
	}))
	if err != nil {
		log.Println(err)
	}
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
		}
	}()
	callForceOrder := gate.NewCallBack(func(msg *gate.UpdateMsg) {
		// parse the message to struct we need
		var order gate.FuturesLiquidate
		if err := sonic.Unmarshal(msg.Result, &order); err != nil {
			log.Printf("order Unmarshal err:%s", err.Error())
		}
		side := "LONG"
		if order.Size < 0 {
			side = "SHORT"
		}
		if order.Contract != "" {
			text := fmt.Sprintf(`
			<b>Gateio</b>
			<b>交易对:%s</b>
			成交额:%d	
			<b>方向:%s</b>
			`, order.Contract, order.Size, side)
			fmt.Println(text)
		}
	})
	ws.SetCallBack(gate.ChannelFutureLiquidates, callForceOrder)
	if err := ws.Subscribe(gate.ChannelFutureLiquidates, []string{"!all"}); err != nil {
		log.Printf("Subscribe err:%s", err.Error())
		return
	}
	for {
	}
}
