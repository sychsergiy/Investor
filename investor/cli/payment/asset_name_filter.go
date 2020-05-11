package payment

import (
	"fmt"
	"investor/entities/payment"
	"investor/interactors/payment_filters"
	"log"
	"strings"
)

type FilterByAssetNamesCommand struct {
	interactor payment_filters.AssetNamesFilter
}

func (c FilterByAssetNamesCommand) Execute() {
	assetNames := readAssetNames()
	paymentTypes := choosePaymentTypes()

	req := payment_filters.AssetNameFilterRequest{
		Periods:      []payment.Period{},
		PaymentTypes: paymentTypes,
		AssetNames:   assetNames,
	}
	resp, err := c.interactor.Filter(req)
	if err != nil {
		log.Fatal(err)
	}
	printPayments(resp.Payments)
}

func NewFilterByAssetNamesCommand(interactor payment_filters.AssetNamesFilter) FilterByAssetNamesCommand {
	return FilterByAssetNamesCommand{interactor: interactor}
}

func readAssetNames() []string {
	fmt.Printf("Enter asset names, use space as delimiter: ")
	input := readFromConsole()
	assetNames := strings.Split(input, " ")
	fmt.Printf("Entered asset names: [%s]\n", strings.Join(assetNames, ", "))
	return assetNames
}
