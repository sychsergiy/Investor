package asset

import (
	"Investor/asset/amount"
	"Investor/asset/amount/crypto"
	"Investor/asset/amount/fiat"
	"Investor/asset/payment"
	"Investor/asset/payment/storage"
	"time"
)

func MakeInvestment(storage storage.PaymentsSaver, invested amount.Amount, cryptoInvestedAmount float32, date time.Time) {
	p := payment.NewInvestPayment(invested.AbsoluteValue(), cryptoInvestedAmount, date)
	storage.SavePayment(p)
}

func MakeReturn(storage storage.PaymentsSaver, returned amount.Amount, cryptoReturnedAmount float32, date time.Time) {
	p := payment.NewReturnPayment(returned.AbsoluteValue(), cryptoReturnedAmount, date)
	storage.SavePayment(p)
}

func CreateInvestment(invested float32, fiatCurrency fiat.Currency, cryptoCurrency crypto.Currency) {
	// choose fiat Currency

	// choose crypto Currency

	// input invested amount

	// suggest USD amount
	// maybe manual entering

	// suggest crypto crypto Amount
	// maybe manual entering

	// input date
}

