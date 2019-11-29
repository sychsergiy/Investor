package storage

import (
	"Investor/asset/payment"
)

type InMemory struct {
	payments map[int]payment.Payment
}

func NewInMemory() *InMemory {
	return &InMemory{payments: make(map[int]payment.Payment)}
}

func getNextID(payments map[int]payment.Payment) int {
	max := -1
	for key := range payments {
		if key > max {
			max = key
		}
	}
	return max + 1
}

func (s *InMemory) SavePayment(payment payment.Payment) {
	s.payments[getNextID(s.payments)] = payment
}

func (s *InMemory) RetrieveAllPayments() (payments []payment.Payment) {
	for _, p := range s.payments {
		payments = append(payments, p)
	}
	return
}

func (s *InMemory) DeletePayment(id int) bool {
	if _, exists := s.payments[id]; exists {
		delete(s.payments, id)
		return true
	}
	return false
}
