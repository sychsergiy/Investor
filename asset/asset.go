package asset

import (
	"Investor/asset/amount/fiat"
	"Investor/asset/period"
)

type ProfitCalculator interface {
	CalcProfit() Profit
	CalcProfitOnPeriod(p period.Period) Profit
}

type CurrencyRateCalculator interface {
	// calculate fiat exchange rate.
	// If investment returned with calculated rate investor will have profit (from arguments)
	// gives possibility to calculate which rate you need to have when returning to get expected profit
	CalcCurrencyRate(profit Profit) fiat.Rate
	CalcCurrencyRateOnPeriod(profit Profit, p period.Period) fiat.Rate
}

type Asset interface {
	ProfitCalculator
	CurrencyRateCalculator
}

type Profit float32

func (p Profit) Percents() Profit {
	return p * 100
}
