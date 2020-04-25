package ports

import (
	"fmt"
	"investor/entities"
)

type InMemoryStorage struct {
	payments map[string]entities.Payment
}

type PaymentAlreadyExitsError struct {
	PaymentId string
}

func (e PaymentAlreadyExitsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.PaymentId)
}

func (storage *InMemoryStorage) Create(payment entities.Payment) (err error) {
	_, idExists := storage.payments[payment.Id]
	if idExists {
		err = PaymentAlreadyExitsError{PaymentId: payment.Id}
	} else {
		storage.payments[payment.Id] = payment
	}
	return
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{payments: make(map[string]entities.Payment)}
}
