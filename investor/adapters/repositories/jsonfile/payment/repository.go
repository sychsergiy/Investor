package payment

import (
	"investor/adapters/repositories/in_memory"
	inMemoryPayment "investor/adapters/repositories/in_memory/payment"
	"investor/adapters/repositories/jsonfile"
	paymentEntity "investor/entities/payment"
)

type JsonFileRepository struct {
	repository jsonfile.Repository
}

func (r JsonFileRepository) CreateBulk(payments []paymentEntity.Payment) (int, error) {
	var records []in_memory.Record
	for _, payment := range payments {
		records = append(records, inMemoryPayment.Record{Payment: payment})
	}
	return r.repository.CreateBulk(records)
}

func (r JsonFileRepository) Create(payment paymentEntity.Payment) error {
	return r.repository.Create(inMemoryPayment.Record{Payment: payment})
}

func NewRepository(repository jsonfile.Repository) JsonFileRepository {
	return JsonFileRepository{repository}
}
