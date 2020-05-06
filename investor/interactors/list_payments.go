package interactors

import paymentEntity "investor/entities/payment"

type ListPayments struct {
	Repository PaymentsLister
}

func (lp ListPayments) ListAll() []paymentEntity.Payment {
	return lp.Repository.ListAll()
}
