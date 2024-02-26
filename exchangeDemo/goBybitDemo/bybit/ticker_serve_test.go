package bybit

import (
	"testing"
)

func TestTickerServe(t *testing.T) {
	clt := newClient()

	bc := clt.NewPublicWsClient(Futures)

	bc.SubTicker([]WsTickerArg{
		{
			Category: Futures,
			Pair:     "ETHUSDT",
		},
		// {
		// 	Category: exchange.Spot,
		// 	Pair:     exchange.MustPair("ETH/USDT"),
		// },
	}, func(k *Ticker) {
		t.Log(*k)
	})

	bc.Start()
}
