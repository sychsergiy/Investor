package crypto

import (
	"Investor/asset"
	"Investor/asset/amount/fiat"
	"Investor/asset/payment"
	"Investor/asset/payment/storage"
	"Investor/asset/period"
)

type CurrencyAsset struct {
	storage.Storage
}

func (a CurrencyAsset) CalcProfit() asset.Profit {
	return CalculateProfit(filterInvestsAndReturns(a.Storage.RetrieveAllPayments()))
}

func (a CurrencyAsset) CalcProfitOnPeriod(period period.Period) asset.Profit {
	payments := a.Storage.RetrieveAllPayments()
	payments = filterPaymentsByPeriod(payments, period)
	invests, returns := filterInvestsAndReturns(payments)
	return CalculateProfit(invests, returns)
}

func filterInvestsAndReturns(payments []payment.Payment) (invests []payment.Payment, returns []payment.Payment) {
	for _, p := range payments {
		switch p.Type {
		case payment.Invest:
			invests = append(invests, p)
		case payment.Return:
			returns = append(returns, p)
		}
	}
	return
}

func calcSums(payments []payment.Payment) (absoluteSum float32, currencySum float32) {
	for _, item := range payments {
		absoluteSum += item.AbsoluteAmount
		currencySum += item.CurrencyAmount
	}
	return
}

func CalculateProfit(invests []payment.Payment, returns []payment.Payment) asset.Profit {
	// calculate asset profit coefficient
	// find invested capital sum(dollars), find invested crypto currency sum
	// a = find percentage of crypto currency rest after all payments
	// b = find percentage of capital was returned
	// profit coefficient: (1 - b / a)
	// profit coefficient 1 means no profit no benefit, 0.5 means lost 50%, 1.5 earned 50%, 5 means earned 400%

	investedSum, cryptoInvestedSum := calcSums(invests)
	if investedSum == 0 {
		panic("No invests made, can't calculate profit") // todo: handle it better

	}
	returnedSum, cryptoReturnedSum := calcSums(returns)

	cryptoRestPart := (cryptoInvestedSum - cryptoReturnedSum) / cryptoInvestedSum
	cryptoSpendPart := 1 - cryptoRestPart
	if cryptoSpendPart == 0 {
		panic("No returns made, can't calculate profit") // todo: handle it better
	}

	returnPart := returnedSum / investedSum

	return asset.Profit(returnPart / cryptoSpendPart)
}

func (a CurrencyAsset) CalcCurrencyRate(profit asset.Profit) fiat.Rate {
	invests, returns := filterInvestsAndReturns(a.Storage.RetrieveAllPayments())
	return calcCurrencyRate(invests, returns, profit)
}

func (a CurrencyAsset) CalcCurrencyRateOnPeriod(profit asset.Profit, p period.Period) fiat.Rate {
	payments := a.Storage.RetrieveAllPayments()
	payments = filterPaymentsByPeriod(payments, p)
	invests, returns := filterInvestsAndReturns(payments)
	return calcCurrencyRate(invests, returns, profit)
}

func calcCurrencyRate(
	invests []payment.Payment,
	returns []payment.Payment,
	profit asset.Profit) fiat.Rate {
	investedSum, cryptoInvestedSum := calcSums(invests)
	if investedSum == 0 {
		panic("invested sum is 0")
	}
	returnedSum, cryptoReturnedSum := calcSums(returns)

	wantToHave := investedSum * float32(profit)
	wantToHaveFromRest := wantToHave - returnedSum
	restCrypto := cryptoInvestedSum - cryptoReturnedSum
	if restCrypto == 0 { // todo: or < 0
		panic("No crypto fiat left, can't get desired profit")
	}
	return fiat.Rate(wantToHaveFromRest / restCrypto)
}

func filterPaymentsByPeriod(payments []payment.Payment, period period.Period) (
	filteredPayments []payment.Payment) {
	for _, p := range payments {
		if period.Contains(p.Date) {
			filteredPayments = append(filteredPayments, p)
		}
	}
	return
}
