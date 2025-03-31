package main

// -ade3252d-3bf826a9-a79d5028-cdgs9k03f3
// -333bc3d5-2301c6ca-4bfc06ca-154b4
import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"

	huobisdk "github.com/huobirdcenter/huobi_golang/pkg/client"
)

const (
	API    = "ade3252d-3bf826a9-a79d5028-cdgs9k03f3"
	SECRET = "333bc3d5-2301c6ca-4bfc06ca-154b4"
)

func main() {
	// signer := new(huobisdk.Signer).Init("secret")

	// result, err := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
	// if err != nil {
	// 	t.Fatalf("unexpected error: %v", err)
	// }
	// Get the list of accounts owned by this API user and print the detail on console
	clientnew := new(huobisdk.AccountClient).Init(API, SECRET, "api.huobi.pro", config.Sign)

	resp2, err := clientnew.GetAccountInfo()
	if err != nil {
		applogger.Error("Get account error: %s", err)
	} else {
		applogger.Info("Get account, count=%d", len(resp2))
		for _, result := range resp2 {
			applogger.Info("account: %+v", result)
		}
	}
}
