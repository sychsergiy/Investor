package payment

import (
	"fmt"
	"investor/entities/payment"
	"investor/interactors"
	"log"
)

type CalcProfitCommand struct {
	profitCalculator interactors.CalcAssetsProfit
}

func (l CalcProfitCommand) Execute() {
	l.CalculateProfit()
}

func (l CalcProfitCommand) CalculateProfit() {
	assetNames := readAssetNames()

	req := interactors.CalcProfitRequest{
		AssetNames: assetNames,
		Periods:    []payment.Period{},
	}
	resp, err := l.profitCalculator.Calc(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nProfit results:")
	for _, p := range resp.Profits {
		p.Profit.Percentage()
		fmt.Printf(
			"Asset Name: %s - %.1f (%.2f%%), payments count: %d;\n",
			p.AssetName, p.Profit.Coefficient(), p.Profit.Percentage(), p.PaymentsCount,
		)
	}
}

func NewCalcProfitCommand(
	profitCalculator interactors.CalcAssetsProfit,
) CalcProfitCommand {
	return CalcProfitCommand{profitCalculator}
}
