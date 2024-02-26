package bybit

import (
	"encoding/json"

	"github.com/dolotech/log"
)

type WsTickerArg struct {
	Category string
	Pair     string
}

type Ticker struct {
	Symbol            string  `json:"symbol"`
	TickDirection     string  `json:"tickDirection"`
	Price24hPcnt      string  `json:"price24hPcnt"`
	LastPrice         float64 `json:"lastPrice,string"`
	PrevPrice24h      float64 `json:"prevPrice24h,string"`
	HighPrice24h      string  `json:"highPrice24h"`
	LowPrice24h       string  `json:"lowPrice24h"`
	PrevPrice1h       string  `json:"prevPrice1h"`
	MarkPrice         string  `json:"markPrice"`
	IndexPrice        string  `json:"indexPrice"`
	OpenInterest      string  `json:"openInterest"`
	OpenInterestValue string  `json:"openInterestValue"`
	Turnover24h       float64 `json:"turnover24h,string"`
	Volume24h         float64 `json:"volume24h,string"`
	NextFundingTime   string  `json:"nextFundingTime"`
	FundingRate       string  `json:"fundingRate"`
	Bid1Price         string  `json:"bid1Price"`
	Bid1Size          string  `json:"bid1Size"`
	Ask1Price         string  `json:"ask1Price"`
	Ask1Size          string  `json:"ask1Size"`
}

func (c *PublicWsClient) SubTicker(args []WsTickerArg, handler func(*Ticker)) error {
	cached := map[string]*Ticker{}
	for _, a := range args {
		a := a
		err := c.sub("tickers."+a.Pair, func(data []byte) {
			// fmt.Println(string(data))
			raw := new(Ticker)
			err := json.Unmarshal(data, raw)
			if err != nil {
				log.Error("kline handler", err, "data", string(data))
				return
			}

			t := new(Ticker)
			t.Symbol = a.Pair

			if raw.LastPrice == 0 {
				t.LastPrice = cached[raw.Symbol].LastPrice
			} else {
				t.LastPrice = raw.LastPrice
			}
			if raw.PrevPrice24h == 0 {
				t.PrevPrice24h = cached[raw.Symbol].PrevPrice24h
			} else {
				t.PrevPrice24h = raw.PrevPrice24h
			}
			if raw.Turnover24h == 0 {
				t.Turnover24h = cached[raw.Symbol].Turnover24h
			} else {
				t.Turnover24h = raw.Turnover24h
			}
			if raw.Volume24h == 0 {
				t.Volume24h = cached[raw.Symbol].Volume24h
			} else {
				t.Volume24h = raw.Volume24h
			}

			// parseKline(t, raw)
			handler(t)
			cached[raw.Symbol] = t
		})
		if err != nil {
			return err
		}
	}
	return nil
}
