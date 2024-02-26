package bybit

import (
	"context"
	"testing"
)

func TestBalanceService(t *testing.T) {
	clt := newClient()
	clt.Debug = true
	s := clt.NewBalanceService()
	s.Auth(key, secret)

	// s.Coin("USDT", "BZZ")
	s.AccountType("CONTRACT")
	resp, err := s.Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	for _, a := range resp {
		t.Log("account >", a.TotalAvailableBalance, a.TotalEquity, a.TotalWalletBalance)
		for _, cb := range a.Coin {
			t.Log("\tcoin > ", cb.Coin, cb.Equity, cb.WalletBalance, cb.UsdValue)
		}
	}
}

func TestUpgradeUTAService(t *testing.T) {
	clt := newClient()
	s := clt.NewUpgradeUTAService()
	s.Auth(key, secret)
	res, err := s.Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", res)
}

func TestConfigService(t *testing.T) {
	clt := newClient()
	s := clt.NewConfigService()
	s.Auth(key, secret)
	resp, err := s.Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", *resp)
}
