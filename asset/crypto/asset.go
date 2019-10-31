package crypto

import (
	"Investor/asset"
	"Investor/asset/payment"
	"Investor/asset/payment/storage"
	"Investor/asset/period"
)

type CurrencyAsset struct {
	storage.Storage
}

func (a CurrencyAsset) CalcProfit() asset.Profit {
	return CalculateProfit(retrievePayments(a.Storage))
}

func (a CurrencyAsset) CalcProfitOnPeriod(period period.Period) asset.Profit {
	invests, returns := retrievePayments(a.Storage)
	filteredInvests, filteredReturns := filterPayments(invests, returns, period)
	return CalculateProfit(filteredInvests, filteredReturns)
}

func retrievePayments(storage storage.Storage) (invests []payment.InvestPayment, returns []payment.ReturnPayment) {
	payments := storage.RetrieveAllPayments()
	for _, p := range payments {
		switch p.Type {
		case payment.Invest:
			invests = append(invests, payment.InvestPayment{Payment: p})
		case payment.Return:
			returns = append(returns, payment.ReturnPayment{Payment: p})
		}
	}
	return
}
func calcInvestsSums(payments []payment.InvestPayment) (absoluteSum float32, currencySum float32) {
	for _, item := range payments {
		absoluteSum += item.AbsoluteAmount
		currencySum += item.CurrencyAmount
	}
	return
}

func calcReturnsSums(payments []payment.ReturnPayment) (absoluteSum float32, currencySum float32) {
	for _, item := range payments {
		absoluteSum += item.AbsoluteAmount
		currencySum += item.CurrencyAmount
	}
	return
}

func CalculateProfit(invests []payment.InvestPayment, returns []payment.ReturnPayment) asset.Profit {
	// calculate asset profit coefficient
	// find invested capital sum(dollars), find invested crypto currency sum
	// a = find percentage of crypto currency rest after all payments
	// b = find percentage of capital was returned
	// profit coefficient: (1 - b / a)
	// profit coefficient 1 means no profit no benefit, 0.5 means lost 50%, 1.5 earned 50%, 5 means earned 400%

	investedSum, cryptoInvestedSum := calcInvestsSums(invests)
	if investedSum == 0 {
		panic("No invests made, can't calculate profit") // todo: handle it better

	}
	returnedSum, cryptoReturnedSum := calcReturnsSums(returns)

	cryptoRestPart := (cryptoInvestedSum - cryptoReturnedSum) / cryptoInvestedSum
	cryptoSpendPart := 1 - cryptoRestPart
	if cryptoSpendPart == 0 {
		panic("No returns made, can't calculate profit") // todo: handle it better
	}

	returnPart := returnedSum / investedSum

	return asset.Profit(returnPart / cryptoSpendPart)
}

func (a CurrencyAsset) CalcCurrencyRate(profit asset.Profit) float32 {
	invests, returns := retrievePayments(a.Storage)
	return calcCurrencyRate(invests, returns, profit)
}

func (a CurrencyAsset) CalcCurrencyRateOnPeriod(profit asset.Profit, p period.Period) float32 {
	invests, returns := retrievePayments(a.Storage)
	invests, returns = filterPayments(invests, returns, p)
	return calcCurrencyRate(invests, returns, profit)
}

func calcCurrencyRate(
	invests []payment.InvestPayment, returns []payment.ReturnPayment, profit asset.Profit) float32 {
	investedSum, cryptoInvestedSum := calcInvestsSums(invests)
	if investedSum == 0 {
		panic("invested sum is 0")
	}
	returnedSum, cryptoReturnedSum := calcReturnsSums(returns)

	wantToHave := investedSum * float32(profit)
	wantToHaveFromRest := wantToHave - returnedSum
	restCrypto := cryptoInvestedSum - cryptoReturnedSum
	if restCrypto == 0 { // todo: or < 0
		panic("No crypto currency left, can't get desired profit")
	}
	return wantToHaveFromRest / restCrypto
}

func filterPayments(invests []payment.InvestPayment, returns []payment.ReturnPayment, period period.Period) (
	filteredInvests []payment.InvestPayment,
	filteredReturns []payment.ReturnPayment) {
	for _, invest := range invests {
		if period.Contains(invest.Date) {
			filteredInvests = append(filteredInvests, invest)
		}
	}
	for _, return_ := range returns {
		if period.Contains(return_.Date) {
			filteredReturns = append(filteredReturns, return_)
		}
	}
	return
}
