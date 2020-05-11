package payment

import (
	"fmt"
	"investor/entities/asset"
	"investor/entities/payment"
	"investor/interactors/payment_filters"
	"log"
	"strconv"
	"strings"
)

type FilterByCategoriesCommand struct {
	interactor payment_filters.AssetCategoriesFilter
}

func NewFilterByCategoriesCommand(interactor payment_filters.AssetCategoriesFilter) FilterByCategoriesCommand {
	return FilterByCategoriesCommand{interactor: interactor}
}

func (c FilterByCategoriesCommand) Execute() {
	categories := chooseCategories()
	paymentTypes := choosePaymentTypes()

	req := payment_filters.AssetCategoriesFilterRequest{
		Periods:         []payment.Period{},
		PaymentTypes:    paymentTypes,
		AssetCategories: categories,
	}

	resp, err := c.interactor.Filter(req)
	if err != nil {
		log.Fatal(err)
	}
	printPayments(resp.Payments)
}

func choosePaymentTypes() []payment.Type {
	fmt.Println("Choose payment type:\n 1 - Invest \n 2 - Return\n 0 - Select all")
	switch input := readFromConsole(); input {
	case "1":
		return []payment.Type{payment.Invest}
	case "2":
		return []payment.Type{payment.Return}
	case "0":
		return []payment.Type{}
	default:
		fmt.Printf("Expected input 0 or 1 or 2, but got %s\n", input)
		fmt.Println("Retrying ...")
		return choosePaymentTypes()
	}
}

func chooseCategories() (categories []asset.Category) {
	fmt.Printf(
		"Available categories:\n%d - %s\n%d - %s\n%d - %s\nall - to select all\n",
		asset.PreciousMetal, asset.PreciousMetal,
		asset.CryptoCurrency, asset.CryptoCurrency,
		asset.Stock, asset.Stock,
	)
	fmt.Printf("Enter numbers to select categories, use space as delimiter: ")
	input := readFromConsole()

	if input == "all" {
		categories = []asset.Category{asset.PreciousMetal, asset.Stock, asset.CryptoCurrency}
	} else {
		categoriesStr := strings.Split(input, " ")
		for _, c := range categoriesStr {
			number, err := strconv.Atoi(c)
			if err != nil {
				fmt.Printf("Unexpected input, number expeted but got %s.\nRetrying ...\n", c)
				return chooseCategories()
			}
			switch number {
			case int(asset.PreciousMetal):
				categories = append(categories, asset.PreciousMetal)
			case int(asset.CryptoCurrency):
				categories = append(categories, asset.CryptoCurrency)
			case int(asset.Stock):
				categories = append(categories, asset.Stock)
			}
		}
	}

	fmt.Printf("Selected categories: [")

	for _, category := range categories {
		fmt.Printf("%s, ", category)
	}
	fmt.Print("]\n")
	return
}
