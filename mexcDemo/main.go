package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// test get Account
	// spot := new(mexcsdk.MexcFutureClient).Init("mx0vgl9MW1A354huGP", "ea5f6fadc57d433d8dbcc1e545c51d49", "https://contract.mexc.com")
	// // params := make(map[string]string, 0)
	// // params["symbol"] = "BTCUSDT"
	// // fmt.Println(spot.MexcClient.URL)
	// result, err := spot.GetSymbolDetail(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(result.Data)
	// for _, v := range result.Data {
	// 	fmt.Println(v.BaseCoin)
	// }
	// for v := range result.Symbols {
	// }
	viper.SetConfigFile("cfg.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	symbols := viper.GetStringMapString("spots")
	fmt.Println(symbols)
	// url := "https://api.mexc.com/api/v3/defaultSymbols"
	// c := resty.New()
	// file, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY, 0o755)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// c.Debug = true
	// resp, _ := c.R().Get(url)
	// // fmt.Println(string(resp.Body()))
	// // res := make([]string, 0)
	// res := Data{}
	// sonic.Unmarshal(resp.Body(), &res)
	// for i := range symbols {
	// 	symbol := strings.ToUpper(i) + "USDT"
	// 	if symbol == "BTCUSDT" {
	// 		fmt.Println(symbol)
	// 	}
	// 	for _, v := range res.Data {
	// 		if symbol == v {
	// 			_, err := file.WriteString(fmt.Sprintf("%s:0\n", v))
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 		}
	// 		// fmt.Println(v)
	// 	}
	// }
	// file.Close()
}

type Data struct {
	Data []string `json:"data"`
}
