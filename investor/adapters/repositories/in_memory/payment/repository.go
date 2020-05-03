package payment

import (
	"investor/adapters/repositories/in_memory"
	paymentEntity "investor/entities/payment"
)

type InMemoryRepository struct {
	repository in_memory.Repository
}

type Record struct {
	paymentEntity.Payment
}

func (p Record) Id() string {
	return p.Payment.Id
}

func (r *InMemoryRepository) Create(payment paymentEntity.Payment) error {
	record := Record{Payment: payment}
	return r.repository.Create(record)
}

func (r *InMemoryRepository) CreateBulk(payments []paymentEntity.Payment) (int, error) {
	var records []in_memory.Record
	for _, payment := range payments {
		records = append(records, Record{Payment: payment})
	}
	return r.repository.CreateBulk(records)
}

func NewRepository() *InMemoryRepository {
	return &InMemoryRepository{in_memory.NewRepository()}
}
