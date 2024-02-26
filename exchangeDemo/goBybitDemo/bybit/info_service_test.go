package bybit

import (
	"context"
	"testing"
)

func newClient() *Client {
	c := NewClient(SetBaseURL("https://api-testnet.bybit.com"))
	// c.Debug = true
	return c
}

func TestInfoService(t *testing.T) {
	clt := newClient()
	s := clt.NewInfoService()
	clt.Debug = true
	// resp, err := s.Category(exchange.Futures).Pair(exchange.MustPair("BTC/USDT")).Do(context.Background())
	resp, err := s.Category(Spot).Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	for _, f := range resp.List {
		t.Logf("%+v", f)
	}
	// for _, f := range resp.Info() {
	// 	// t.Log(f)
	// 	t.Log(f.Step, f.Value, f.MinSize)
	// }
}
