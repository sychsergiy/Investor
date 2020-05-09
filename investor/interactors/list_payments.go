package interactors

import (
	"investor/entities/payment"
	"investor/interactors/ports"
)

type ListPayments struct {
	Repository ports.PaymentsLister
}

func (lp ListPayments) ListAll() ([]payment.Payment, error) {
	return lp.Repository.ListAll()
}
