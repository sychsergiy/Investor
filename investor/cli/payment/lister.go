package payment

import (
	"fmt"
	"investor/entities/payment"
	"investor/interactors"
	"log"
)

type ListPaymentsCommand struct {
	lister interactors.ListPayments
}

func (l ListPaymentsCommand) Execute() {
	l.List()
}

func (l ListPaymentsCommand) List() {
	payments, err := l.lister.ListAll()
	if err != nil {
		log.Fatalf("failed to list payments: %+v", err)
	}
	printPayments(payments)
}

func printPayments(payments []payment.Payment) {
	fmt.Printf("----------Total payments count: %d----------\n", len(payments))
	for i, p := range payments {
		str, err := paymentToString(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%d -------------------------\n", i+1)
		println(str)
	}
	fmt.Printf("-------------------END---------------------\n\n")
}

func paymentToString(p payment.Payment) (str string, err error) {
	asset, err := p.Asset()
	if err != nil {
		return
	}
	str = fmt.Sprintf(
		"ID: %s\nAsset name: %s\nAsset category: %s\nType: %s\nUSD amount: %f\nAsset amount: %f\nCreation date: %s",
		p.Id(), asset.Name(), asset.Category().String(), p.Type().String(), p.AbsoluteAmount(),
		p.AssetAmount(), p.CreationDate().Format("2006-01-02 15:04"),
	)
	return

}

func NewConsolePaymentsLister(paymentsLister interactors.ListPayments) ListPaymentsCommand {
	return ListPaymentsCommand{paymentsLister}
}
