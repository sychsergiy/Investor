package in_memory

import (
	"investor/adapters/repositories"
	paymentEntity "investor/entities/payment"
)

type InMemoryPaymentRepository struct {
	repository repositories.InMemoryRepository
}

type PaymentRecord struct {
	Payment paymentEntity.Payment
}

func (p PaymentRecord) Id() string {
	return p.Payment.Id
}

func (r *InMemoryPaymentRepository) Create(payment paymentEntity.Payment) error {
	record := PaymentRecord{Payment: payment}
	return r.repository.Create(record)
}

func (r *InMemoryPaymentRepository) CreateBulk(payments []paymentEntity.Payment) (int, error) {
	var records []repositories.Record
	for _, payment := range payments {
		records = append(records, PaymentRecord{Payment: payment})
	}
	return r.repository.CreateBulk(records)
}

func NewInMemoryPaymentRepository(repository repositories.InMemoryRepository) *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{repository: repository}
}
