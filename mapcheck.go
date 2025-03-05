package main

import "fmt"

var runningState = map[uint8]map[string]string{
	0: {"Waiting Signal": "The bot is in the monitoring state, and the position will be opened after a short while."},
	1: {"Running": ""},
	2: {"Smaller Position": `The position size recorded by the bot is inconsistent with the actual position size of the exchange. The reason is that the user manually reduced or closed all positions from the exchange. 
<b>Suggestion: Restart the bot after performing the “Stop” operation on the bot.</b>`},
	3: {"Smaller Position": ""},
	4: {"Insufficient USDT": `it means that your USDT in the account of the exchange is insufficient, the robot will continue to maintain the current position and will can not open new position.
<b>Suggestion: Please go to the exchange to recharge USDT, or wait for the bots are profitable and close position automatically to free up the available funds.</b>`},
	5: {"Insufficient Margin": `It means that your margin balance in the exchange futures account is insufficient.
<b>Suggestion: Please promptly go to the exchange to increase your margin to avoid the risk of liquidation.</b>`},
	6: {"Unallocated Funds": `the robot of opening positions is constrained by the "Invest Control". If continued to open positions, it will exceed the preset limit by the "Invest Control" settings.
<b>Suggestion: you wish to resume the robot's operations of opening positions, you can add margin on the exchange or or go to the [Home] - [Invest Control] page to adjust the relevant risk control parameters.</b>`},
	7: {"Price Ceiling": ""},
	8: {"Max Positions": `It means that the current robot has reached its maximum position limit, and the position control system prevents further to covering the position.
<b>Suggestion: You can patiently wait for market changes. or if you want to continue to covering the position with the robot to bring the average position price closer to the current market price, it is recommended to add USDT to the futures account on the exchange. This will enable you to generate profits earlier and create more opportunities for profit.</b>`},
	9: {"Limit Order": `The number of orders has been used up, and there will be no more replenishment. 
<b>Suggestion: Users are required to judge whether they need to increase the number of orders based on their available funds and market conditions.</b>`},
	10: {"Trading Limit": `When this state occurs, it is generally because the current position is close to the maximum position limit of the exchange, and the position can no longer be covered, but the position can only be closed. 
<b>Suggestion: You can wait for the bots are profitable and close position automatically.</b>`},
	11: {"Trading Error": `This status cannot be automatically repaired. Generally, the exchange has rejected the transaction request and the robot cannot trade normally. Please contact uTrading online support for help.
<b>Suggestion: Check whether the API is normal. If the API is not normal, you need to go to exchange to get the new API Key. Import uTrading again.</b>`},
	12: {"Price Floor": `When the currency falls below the “Price Floor”, the bot will no longer open or cover positions. 
<b>Suggestion: The bot that supports the AI strategy will automatically adjust the price floor.
Bots with non-AI strategies can manually modify the “Price Floor”.</b>`},
	13: {"Waiting Covering": "When this state occurs, the position replenishment will be suspended, and the position will be replenished after waiting for the general trend to change."},
	14: {"Min position less than 15u": "Indicates that the current initial positions amount or the investment amount is less than the transaction amount."},
	15: {"Waiting Trend": "The robot meets the conditions for opening positions. it will wait for long or short signals to appear, then choose a trading direction to open positions."},
	16: {"API Error": ""},
	17: {"Waiting Directions": "The robot will waiting for signals either long or short positions, then select a trading direction and open positions."},
	18: {"Country Limiting": ""},
	19: {"Leverage Limit": ""},
	20: {"Trading Capital Limit": ""},
	21: {"Waiting Conditions": ""},
	22: {"Position Mode Error": ""},
	23: {"Price Ceiling": ""},
	24: {"Sync Limit": `The allocated Follow Investment has been used, and there are no remaining funds for the sync robot to open or add positions.
<b>Suggestion: Adjust the Follow Investment Limit to increase your investment, Or wait patiently, the robot holds existing positions to run.</b>`},
}

func main() {
	// key, value := runningState[6]

	for key, value := range runningState[6] {
		fmt.Println(key)
		fmt.Println(value)
	}
}
