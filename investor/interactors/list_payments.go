package interactors

import paymentEntity "investor/entities/payment"

type ListPayments struct {
	Repository PaymentsLister
}

func (lp ListPayments) ListAll() ([]paymentEntity.Payment, error) {
	return lp.Repository.ListAll()
}
