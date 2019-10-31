package asset

import (
	"Investor/asset/period"
)

type Asset interface {
	CalcProfit() Profit
	CalcProfitOnPeriod(p period.Period) Profit

	// calculate currency exchange rate.
	// If investment returned with calculated rate investor will have profit (from arguments)
	// gives possibility to calculate which rate you need to have when returning to get expected profit
	CalcCurrencyRate(profit Profit) float32 // todo: type for currency exchange rate
	CalcCurrencyRateOnPeriod(profit Profit, p period.Period) float32
}

type Profit float32

func (p Profit) Percents() Profit {
	return p * 100
}
