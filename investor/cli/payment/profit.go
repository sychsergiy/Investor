package payment

import (
	"fmt"
	"investor/interactors"
	"log"
)

type CalcProfitCommand struct {
	lister           interactors.ListPayments
	profitCalculator interactors.CalcProfit
}

func (l CalcProfitCommand) Execute() {
	l.CalculateProfit()
}

func (l CalcProfitCommand) CalculateProfit() {
	payments, err := l.lister.ListAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Calculate profit for the following payments:")
	printPayments(payments)
	profit, err := l.profitCalculator.Calc(payments)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Profit coeficient: %f\nProfit percentage: %f", profit.Coefficient(), profit.Percentage())
}

func NewCalcProfitCommand(
	paymentsLister interactors.ListPayments,
	profitCalculator interactors.CalcProfit,
) CalcProfitCommand {
	return CalcProfitCommand{paymentsLister, profitCalculator}
}
