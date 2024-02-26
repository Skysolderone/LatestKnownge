package bybit

import (
	"context"
	"testing"
)

func TestPositionService(t *testing.T) {
	clt := newClient()
	ps := clt.NewPositionService()
	ps.Auth(key, secret)
	ps.Category(Futures)
	ps.SettleCoin("USDT")

	res, err := ps.Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)
}

func TestChangePositionModeService(t *testing.T) {
	clt := newClient()
	os := clt.NewChangePositionModeService()
	os.Auth(key, secret)
	err := os.Category(Futures).SettleCoin("USDT").Mode(true).Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
}

func TestChangeLeverageService(t *testing.T) {
	clt := newClient()

	ls := clt.NewChangeLeverageService()
	ls.Auth(key, secret)

	err := ls.Category(Futures).Pair("BTCUSDT").Leverage(20).Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
}

func TestChangeMarginModeService(t *testing.T) {
	clt := newClient()
	os := clt.NewChangeMarginModeService()
	os.Auth(key, secret)
	err := os.
		Category(Futures).
		Pair("ETHUSDT").
		TradeMode(true).
		Leverage(10).
		Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
}
