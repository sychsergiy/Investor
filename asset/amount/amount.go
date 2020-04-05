package amount

import (
	"Investor/asset/amount/fiat"
	"fmt"
)

type Amount interface {
	AbsoluteValue() float32 // absolute value for all currencies - $dollars
}

type USD struct {
	Value float32
}

func (a *USD) AbsoluteValue() float32 {
	return convertToUSD(fiat.USD, a.Value)
}

type UAH struct {
	Value float32
}

func (a *UAH) AbsoluteValue() float32 {
	return convertToUSD(fiat.UAH, a.Value)
}

func convertToUSD(inCur fiat.Currency, amount float32) float32 {
	switch inCur {
	case fiat.USD:
		return amount
	case fiat.UAH:
		return amount / 25
	}
	panic(fmt.Sprintf("Unresolved fiat: %s", inCur))
}
