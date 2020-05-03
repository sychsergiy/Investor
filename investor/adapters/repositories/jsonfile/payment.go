package jsonfile

import (
	"investor/adapters/repositories/in_memory"
	paymentEntity "investor/entities/payment"
)

type PaymentRepository struct {
	repository Repository
}

func (r PaymentRepository) CreateBulk(payments []paymentEntity.Payment) (int, error) {
	var records []in_memory.Record
	for _, payment := range payments {
		records = append(records, in_memory.PaymentRecord{Payment: payment})
	}
	return r.repository.CreateBulk(records)
}

func (r PaymentRepository) Create(payment paymentEntity.Payment) error {
	return r.repository.Create(in_memory.PaymentRecord{Payment: payment})
}

func NewPaymentRepository(repository Repository) PaymentRepository {
	return PaymentRepository{repository}
}
