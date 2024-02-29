package main

import "fmt"

type SigtradeSetting struct {
	Status    uint8 `json:"status"`    // 1 启用，2 停止
	Duration  int64 `json:"duration"`  // 运行期限，单位：小时
	Countdown int64 `json:"countdown"` // 倒计时，单位：秒
	Robots    int64 `json:"robots"`    // 配置机器人数量
}
type GetSigtradeJsoned struct {
	Config struct {
		MaxLeverage        uint8   `json:"maxLeverage"`
		MaxInitialPosition float64 `json:"maxInitialPosition"`

		MaxSpotRobots        int64 `json:"maxSpotRobots"`        // 最大机器人
		RunningSpotRobots    int64 `json:"runningSpotRobots"`    // 运行机器人数量
		MaxFuturesRobots     int64 `json:"maxFuturesRobots"`     // 最大机器人
		RunningFuturesRobots int64 `json:"runningFuturesRobots"` // 运行机器人数量
	} `json:"config"`
	SetSigtradeJson
}
type SetSigtradeJson struct {
	Leverage        *uint8   `json:"leverage"`
	InitialPosition *float64 `json:"initialPosition"`

	Spot    *SigtradeSetting `json:"spot"`
	Futures *SigtradeSetting `json:"futures"`
}

func main() {
	for i := range 5 {
		fmt.Println(i)
	}
	res := GetSigtradeJsoned{}
	res.Spot = &SigtradeSetting{
		Duration: 55,
	}
	fmt.Println(res.Spot)
}
