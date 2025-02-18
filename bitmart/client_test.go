package bitmart

import "testing"

func newSpotbitmart() *Spot {
	api := "39237ed9824090889dfc875e39d41a47dc56bf49"
	sec := "83fa4361a186a7e3e074c8dd9c755dd51ac016e1ff86ff628b01ba6e25695970"
	ProxyURL := "https://api-cloud.bitmart.com"
	Memo := "uTrading"
	return NewSpot(api, sec, ProxyURL, Memo)
}

func TestGetOrderbitmart(t *testing.T) {
	// 896010099772304640 buy
	// 697920774684 sell
	c := newSpotbitmart()
	c.GetOrder("896046636051021312")
}

func TestGetBalanc(t *testing.T) {
	c := newSpotbitmart()
	c.Balance()
}
