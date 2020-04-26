package in_memory

import (
	"fmt"
	paymentEntity "investor/entities/payment"
)

type InMemoryPaymentRepository struct {
	payments map[string]paymentEntity.Payment
}

type PaymentAlreadyExistsError struct {
	PaymentId string
}

func (e PaymentAlreadyExistsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.PaymentId)
}

func (storage *InMemoryPaymentRepository) Create(payment paymentEntity.Payment) (err error) {
	_, idExists := storage.payments[payment.Id]
	if idExists {
		err = PaymentAlreadyExistsError{PaymentId: payment.Id}
	} else {
		storage.payments[payment.Id] = payment
	}
	return
}

func (storage *InMemoryPaymentRepository) CreateBulk(payments []paymentEntity.Payment) (int, error) {
	var createdCount int
	for createdCount, payment := range payments {
		_, idExists := storage.payments[payment.Id]
		if idExists {
			return createdCount, PaymentAlreadyExistsError{PaymentId: payment.Id}
		} else {
			storage.payments[payment.Id] = payment
		}
	}
	return createdCount, nil
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{payments: make(map[string]paymentEntity.Payment)}
}
