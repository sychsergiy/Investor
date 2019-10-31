package amount

import (
	"Investor/asset/amount/currency"
	"fmt"
)

type Amount interface {
	AbsoluteValue() float32 // absolute value for all currencies - $dollars
}

type USD struct {
	Value float32
}

func (a *USD) AbsoluteValue() float32 {
	return convertToUSD(currency.USD, a.Value)
}

type UAH struct {
	Value float32
}

func (a *UAH) AbsoluteValue() float32 {
	return convertToUSD(currency.UAH, a.Value)
}

func convertToUSD(inCur currency.Currency, amount float32) float32 {
	switch inCur {
	case currency.USD:
		return amount
	case currency.UAH:
		return amount / 25
	}
	panic(fmt.Sprintf("Unresolved currency: %s", inCur))
}
