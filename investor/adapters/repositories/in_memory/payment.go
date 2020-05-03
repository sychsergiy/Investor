package in_memory

import (
	paymentEntity "investor/entities/payment"
)

type PaymentRepository struct {
	repository Repository
}

type PaymentRecord struct {
	paymentEntity.Payment
}

func (p PaymentRecord) Id() string {
	return p.Payment.Id
}

func (r *PaymentRepository) Create(payment paymentEntity.Payment) error {
	record := PaymentRecord{Payment: payment}
	return r.repository.Create(record)
}

func (r *PaymentRepository) CreateBulk(payments []paymentEntity.Payment) (int, error) {
	var records []Record
	for _, payment := range payments {
		records = append(records, PaymentRecord{Payment: payment})
	}
	return r.repository.CreateBulk(records)
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{NewRepository()}
}
