package interactors

import (
	"investor/entities/payment"
	"investor/interactors/ports"
)

type ListPayments struct {
	repository ports.PaymentsLister
}

func (lp ListPayments) ListAll() ([]payment.Payment, error) {
	return lp.repository.ListAll()
}

func NewListPayments(repository ports.PaymentsLister) ListPayments {
	return ListPayments{repository: repository}
}
