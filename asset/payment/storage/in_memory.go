package storage

import "Investor/asset/payment"

type InMemory struct {
	payments []payment.Payment
}

func NewInMemory() *InMemory {
	return &InMemory{payments: []payment.Payment{}}
}

func (s *InMemory) AddPayment(payment payment.Payment) {
	s.payments = append(s.payments, payment)
}

func (s *InMemory) RetrieveAllPayments() []payment.Payment {
	return s.payments
}
