package payment

import (
	"fmt"
	"investor/entities/payment"
	"investor/interactors"
	"investor/interactors/payment_filters"
	"log"
)

type CalcProfitCommand struct {
	filter           payment_filters.AssetNamesFilter
	profitCalculator interactors.CalcProfit
}

func (l CalcProfitCommand) Execute() {
	l.CalculateProfit()
}

func (l CalcProfitCommand) CalculateProfit() {
	assetNames := readAssetNames()

	var model = payment_filters.AssetNameFilterRequest{
		Periods:      []payment.Period{},
		PaymentTypes: []payment.Type{},
		AssetNames:   assetNames,
	}
	resp, err := l.filter.Filter(model)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Calculate profit for the following payments:")
	printPayments(resp.Payments)
	profit, err := l.profitCalculator.Calc(resp.Payments)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Profit coeficient: %f\nProfit percentage: %f", profit.Coefficient(), profit.Percentage())
}

func NewCalcProfitCommand(
	paymentsFilter payment_filters.AssetNamesFilter,
	profitCalculator interactors.CalcProfit,
) CalcProfitCommand {
	return CalcProfitCommand{paymentsFilter, profitCalculator}
}
