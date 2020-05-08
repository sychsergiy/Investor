package payment

import (
	"fmt"
	"investor/entities/payment"
	"investor/interactors"
)

type ConsolePaymentsLister struct {
	lister interactors.ListPayments
}

func (l ConsolePaymentsLister) List() {
	payments, err := l.lister.ListAll()
	if err != nil {
		panic(fmt.Errorf("failed to list payments: %+v", err))
	}

	fmt.Printf("Total payments count: %d\n", len(payments))
	for i, p := range payments {
		str, err := paymentToString(p)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n%d -------------------------\n", i+1)
		println(str)
	}
}

func paymentToString(p payment.Payment) (str string, err error) {
	asset, err := p.Asset()
	if err != nil {
		return
	}
	str = fmt.Sprintf(
		"ID: %s\nAsset: %s\nAsset category: %s\nType: %s\nUSD amount: %f\nAsset amount: %f\nCreation date: %s",
		p.Id(), asset.Name(), asset.Category().String(), p.Type().String(), p.AbsoluteAmount(),
		p.AssetAmount(), p.CreationDate().Format("2006-01-02 15:04"),
	)
	return

}

func NewConsolePaymentsLister(paymentsLister interactors.ListPayments) ConsolePaymentsLister {
	return ConsolePaymentsLister{paymentsLister}
}
