package bybit

import (
	"context"
	"testing"
)

func TestKlineService(t *testing.T) {
	clt := newClient()
	s := clt.NewKlineService()

	resp, err := s.
		Pair("BTCUSDT").
		Interval(900).
		Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	for i, k := range resp {
		t.Log(i, k)
	}
}
