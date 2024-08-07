package main

import (
	"fmt"
	"time"

	"github.com/nanmu42/etherscan-api"
)

func main() {
	// 创建连接指定网络的客户端
	// client := etherscan.New(etherscan.Mainnet, "[your API key]")

	// 或者，如果你要调用的是EtherScan家族的BscScan：
	//
	client := etherscan.NewCustomized(etherscan.Customization{
		Timeout: 15 * time.Second,
		Key:     "X41GPMHEIHVUJDZI2E1YZXHH2X6IVB24W4",
		BaseURL: "https://api-testnet.bscscan.com/api?",
		Verbose: false,
	})

	// （可选）按需注册钩子函数，例如用于速率控制
	// client.BeforeRequest = func(module, action string, param map[string]interface{}) error {
	// 	// ...
	// }
	// client.AfterRequest = func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error) {
	// 	// ...
	// }

	// 查询账户以太坊余额
	balance, err := client.AccountBalance("0x35A42428a5446E35158b90D6339F8eAaEf95c272")
	if err != nil {
		panic(err)
	}
	// 余额以 *big.Int 的类型呈现，单位为 wei
	fmt.Println(balance.Int())

	// 查询token余额
	tokenBalance, err := client.TokenBalance("contractAddress", "holderAddress")
	fmt.Printf("%#v", tokenBalance)
	// 查询出入指定地址的ERC20转账列表
	// transfers, err := client.ERC20Transfers("contractAddress", "address", startBlock, endBlock, page, offset)
	// fmt.Printf("%#v", transfers)c
}
