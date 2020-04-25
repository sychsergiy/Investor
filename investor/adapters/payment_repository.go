package adapters

import (
	"fmt"
	"investor/entities/payment"
)

type InMemoryPaymentRepository struct {
	payments map[string]payment.Payment
}

type PaymentAlreadyExitsError struct {
	PaymentId string
}

func (e PaymentAlreadyExitsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.PaymentId)
}

func (storage *InMemoryPaymentRepository) Create(payment payment.Payment) (err error) {
	_, idExists := storage.payments[payment.Id]
	if idExists {
		err = PaymentAlreadyExitsError{PaymentId: payment.Id}
	} else {
		storage.payments[payment.Id] = payment
	}
	return
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{payments: make(map[string]payment.Payment)}
}
