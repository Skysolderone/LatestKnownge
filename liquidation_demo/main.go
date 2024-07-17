package main

import (
	"fmt"
)

// Account represents a trading account
type Account struct {
	Balance float64 // 账户余额
}

// Position represents a trading position
type Position struct {
	Quantity   float64 // 持仓数量
	EntryPrice float64 // 开仓价格
	Leverage   float64 // 杠杆率
}

// CalculateMaintenanceMarginRatio calculates the maintenance margin ratio based on leverage
func CalculateMaintenanceMarginRatio(leverage float64) float64 {
	// 示例维持保证金比率，根据杠杆率调整
	// 实际情况中，应该根据币安交易所的具体规则进行调整
	switch {
	case leverage <= 10:
		return 0.01 // 1%
	case leverage <= 20:
		return 0.025 // 2%
	case leverage <= 50:
		return 0.05 // 5%
	case leverage <= 100:
		return 0.1 // 10%
	default:
		return 0.2 // 20%
	}
}

// CalculateLiquidationPrice calculates the liquidation price
func CalculateLiquidationPrice(account Account, position Position) float64 {
	initialMargin := position.EntryPrice * position.Quantity / position.Leverage
	maintenanceMarginRatio := CalculateMaintenanceMarginRatio(position.Leverage)
	maintenanceMargin := maintenanceMarginRatio * position.EntryPrice * position.Quantity

	// Liquidation price formula
	liquidationPrice := (position.EntryPrice*position.Quantity - initialMargin + maintenanceMargin) / position.Quantity
	return liquidationPrice
}

func main() {
	// Example parameters
	account := Account{
		Balance: 1346, // 账户余额
	}
	position := Position{
		Quantity:   13.1,   // 持仓数量 (1 BTC)
		EntryPrice: 38.178, // 开仓价格 ($10,000)
		Leverage:   20,     // 杠杆率 (20x)
	}

	liquidationPrice := CalculateLiquidationPrice(account, position)
	fmt.Printf("Liquidation Price: $%.2f\n", liquidationPrice)
}
