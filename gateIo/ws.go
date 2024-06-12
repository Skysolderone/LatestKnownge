package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bytedance/sonic"
	gate "github.com/gateio/gatews/go"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("final.yaml")
	viper.ReadInConfig()
	data := viper.GetStringMapString("spot")
	symbols := make([]string, 0)

	for i := range data {

		fmt.Println(strings.ToUpper(i) + "_USDT")
		symbols = append(symbols, strings.ToUpper(i)+"_USDT")
		// if len(symbols) > 30 {
		// 	break
		// }
	}
	fmt.Println(symbols)
	ws, err := gate.NewWsService(nil, nil, gate.NewConnConfFromOption(&gate.ConfOptions{
		URL: gate.FuturesUsdtUrl,
	}))
	// ws.GetConnConf().URL = gate.FuturesUsdtUrl
	if err != nil {
		log.Printf("NewWsService err:%s", err.Error())
		return
	}

	// go func() {
	// 	ticker := time.NewTicker(time.Second)
	// 	for {
	// 		<-ticker.C
	// 		log.Println("connetion status:", ws.Status())
	// 	}
	// }()

	// callTicker := gate.NewCallBack(func(msg *gate.UpdateMsg) {
	// 	// fmt.Println(string(msg.Error.Message))
	// 	var ticker gate.SpotTickerMsg
	// 	if err := sonic.Unmarshal(msg.Result, &ticker); err != nil {
	// 		log.Printf("trade Unmarshal err:%s", err.Error())
	// 	}
	// 	log.Printf("%+v", ticker)
	// })

	// ws.SetCallBack(gate.ChannelSpotTicker, callTicker)

	// if err := ws.Subscribe(gate.ChannelSpotTicker, symbols); err != nil {
	// 	log.Printf("Subscribe err:%s", err.Error())
	// 	return
	// }
	callTicker := gate.NewCallBack(func(msg *gate.UpdateMsg) {
		// fmt.Println(string(msg.Error.Message))
		var ticker []gate.FuturesTicker
		if err := sonic.Unmarshal(msg.Result, &ticker); err != nil {
			log.Printf("trade Unmarshal err:%s", err.Error())
		}
		log.Printf("%+v", ticker)
	})

	ws.SetCallBack(gate.ChannelFutureTicker, callTicker)

	if err := ws.Subscribe(gate.ChannelFutureTicker, symbols); err != nil {
		log.Printf("Subscribe err:%s", err.Error())
		return
	}
	ch := make(chan os.Signal)
	signal.Ignore(syscall.SIGPIPE, syscall.SIGALRM)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT, syscall.SIGKILL)
	<-ch
}
