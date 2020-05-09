package payment

import (
	"fmt"
	"investor/interactors"
	"log"
)

type ConsoleProfitCalculator struct {
	lister           interactors.ListPayments
	profitCalculator interactors.CalcProfit
}

func (l ConsoleProfitCalculator) Execute() {
	l.CalculateProfit()
}

func (l ConsoleProfitCalculator) CalculateProfit() {
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

func NewConsoleProfitCalculator(
	paymentsLister interactors.ListPayments,
	profitCalculator interactors.CalcProfit,
) ConsoleProfitCalculator {
	return ConsoleProfitCalculator{paymentsLister, profitCalculator}
}
