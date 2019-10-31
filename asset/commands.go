package asset

import (
	"Investor/asset/amount"
	"Investor/asset/payment"
	"Investor/asset/payment/storage"
	"time"
)

func MakeInvestment(storage storage.Storage, invested amount.Amount, cryptoInvestedAmount float32, date time.Time) {
	invest := payment.NewInvestPayment(invested.AbsoluteValue(), cryptoInvestedAmount, date)
	storage.AddPayment(invest.Payment)
}

func MakeReturn(storage storage.Storage, returned amount.Amount, cryptoReturnedAmount float32, date time.Time) {
	return_ := payment.NewReturnPayment(returned.AbsoluteValue(), cryptoReturnedAmount, date)
	storage.AddPayment(return_.Payment)
}
