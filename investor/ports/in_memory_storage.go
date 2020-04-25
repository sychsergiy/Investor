package ports

import (
	"fmt"
	"investor/entities"
)

type InMemoryStorage struct {
	payments map[string]entities.Payment
}

func (storage *InMemoryStorage) Create(payment entities.Payment) (err error) {
	_, idExists := storage.payments[payment.Id]
	if idExists {
		err = fmt.Errorf("payment with id %s already exists", payment.Id)
	} else {
		storage.payments[payment.Id] = payment
	}
	return
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{payments: make(map[string]entities.Payment)}
}
