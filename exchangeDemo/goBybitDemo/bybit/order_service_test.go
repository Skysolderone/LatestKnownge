package bybit

import (
	"context"
	"testing"
)

const (
	// key    = "2y8peVyrIJKEDmJ9Of"
	// secret = "kTnW5hpim5hXgzAz7RgjttU2E3W8Qno2ihmX"
	// key    = "oyvm2uXY8tAS6jANnO"
	// secret = "aydVq3fac96Z3pqcDhiUKoFTwUhSAiMbw2kM"

	// testnet
	// key    = "lfhd8RT130WJ0zLssS"
	// secret = "L4iuGeY24TXYgSMDMXABwwWiQAc7e3WKwQ45"
	key    = "kOUGrl7aXydTgduTFA"
	secret = "x0SEWtRFHtPn7qqNIS5vlTJ99IoUM7HxQcKc"
)

func TestOrderService(t *testing.T) {
	clt := newClient()
	s := clt.NewOrderService()
	s.Auth(key, secret)

	c := Spot
	p := "BTCUSDT"
	s.Pair(c, p)
	s.Side("Buy")
	s.Type("Market")
	s.PosSide("Long")
	s.TimeInForce("FOK")
	s.Size("10")

	resp, err := s.Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", *resp)

	os := clt.NewQueryOrderService()
	os.Auth(key, secret)
	os.Pair(c, p)
	os.OrderID(resp.OrderId)
	qr, err := os.Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", *qr)
}

func TestQueryOrderService(t *testing.T) {
	clt := newClient()
	clt.Debug = true
	os := clt.NewQueryOrderService()
	os.Auth(key, secret)
	os.Pair(Spot, "BTCUSDT")
	os.OrderID("1491933590968197376")
	qr, err := os.Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", *qr)
}
