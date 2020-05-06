package in_memory

import (
	"fmt"
	paymentEntity "investor/entities/payment"
)

type PaymentAlreadyExistsError struct {
	PaymentId string
}

func (e PaymentAlreadyExistsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.PaymentId)

}

type PaymentRepository struct {
	records map[string]paymentEntity.Payment
}

func (r *PaymentRepository) Create(payment paymentEntity.Payment) error {
	_, idExists := r.records[payment.Id]
	if idExists {
		return PaymentAlreadyExistsError{PaymentId: payment.Id}
	} else {
		r.records[payment.Id] = payment
		return nil
	}
}

func (r *PaymentRepository) CreateBulk(payments []paymentEntity.Payment) (int, error) {
	var createdCount int
	for createdCount, payment := range payments {
		_, idExists := r.records[payment.Id]
		if idExists {
			return createdCount, RecordAlreadyExistsError{RecordId: payment.Id}
		} else {
			r.records[payment.Id] = payment
		}
	}
	return createdCount, nil
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{make(map[string]paymentEntity.Payment)}
}
