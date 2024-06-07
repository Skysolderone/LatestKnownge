package main

import (
	"fmt"
	"log"

	"v1/mexcsdk"
)

func main() {
	// test get Account
	spot := new(mexcsdk.MexcFutureClient).Init("mx0vgl9MW1A354huGP", "ea5f6fadc57d433d8dbcc1e545c51d49", "https://contract.mexc.com")
	// params := make(map[string]string, 0)
	// params["symbol"] = "BTCUSDT"
	result, err := spot.GetSymbolDetail(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// for v := range result.Symbols {
	// }
}
