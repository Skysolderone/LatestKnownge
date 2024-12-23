package main

import (
	"fmt"
	"sort"
)

type PlateLeader struct {
	Symbol      string  `json:"symbol"`      // 币种
	PriceChange float64 `json:"priceChange"` // 涨幅度
	Plate       int     `json:"plate"`       // 板块
}

func main() {
	// [{XRPUSDT -10.521 22} {BTCUSDT -1.527 22} {ETHUSDT -4.525 22}]
	plateLeader := []PlateLeader{
		{Symbol: "XRPUSDT", PriceChange: -10.521, Plate: 22},
		{Symbol: "BTCUSDT", PriceChange: -1.527, Plate: 22},
		{Symbol: "ETHUSDT", PriceChange: -4.525, Plate: 22},
	}
	sort.Slice(plateLeader, func(i, j int) bool {
		// fmt.Println(plateLeader[i].PriceChange, plateLeader[j].PriceChange)
		return plateLeader[i].PriceChange > plateLeader[j].PriceChange
	})
	fmt.Println(plateLeader[0])
}
