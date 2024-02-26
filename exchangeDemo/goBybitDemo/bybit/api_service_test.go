package bybit

import (
	"context"
	"testing"
)

func TestApiService(t *testing.T) {
	clt := newClient()
	clt.Debug = true
	s := clt.NewApiService()
	s.Auth(key, secret)

	res, err := s.Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", *res)
}
