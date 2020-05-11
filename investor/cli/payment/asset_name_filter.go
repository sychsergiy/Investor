package payment

import (
	"fmt"
	"investor/entities/payment"
	"investor/interactors/payment_filters"
	"log"
)

type FilterByAssetNameCommand struct {
	interactor payment_filters.AssetNameFilter
}

func (c FilterByAssetNameCommand) Execute() {
	assetName := readAssetName()
	req := payment_filters.AssetNameFilterRequest{
		TimeFrom:  payment.CreateYearDate(2018),
		TimeUntil: payment.CreateYearDate(2021),
		AssetName: assetName,
	}
	resp, err := c.interactor.Filter(req)
	if err != nil {
		log.Fatal(err)
	}
	printPayments(resp.Payments)
}

func NewFilterByAssetNameCommand(interactor payment_filters.AssetNameFilter) FilterByAssetNameCommand {
	return FilterByAssetNameCommand{interactor: interactor}
}

func readAssetName() string {
	fmt.Println("Enter asset name: ")
	assetName := readFromConsole()
	return assetName
}
