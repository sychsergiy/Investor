package interactors

import (
	"investor/entities/payment"
)

func CalcSumsForPayments(payments []payment.Payment) payment.Sums {
	s := payment.Sums{}
	for _, item := range payments {
		switch item.Type() {
		case payment.Return:
			s.Returned += item.AbsoluteAmount()
			s.ReturnedAsset += item.AssetAmount()
		case payment.Invest:
			s.Invested += item.AbsoluteAmount()
			s.InvestedAsset += item.AssetAmount()
		default:
			panic("unexpected payment type")
		}
	}
	return s
}

type CalcProfit struct{}

func (cp CalcProfit) Calc(payments []payment.Payment) (payment.Profit, error) {
	sums := CalcSumsForPayments(payments)
	return payment.CalcProfit(sums)
}

func NewCalcProfit() CalcProfit {
	return CalcProfit{}
}
