package asset

import (
	"investor/adapters/repositories/in_memory"
	assetEntity "investor/entities/asset"
)

type Record struct {
	assetEntity.Asset
}

func (a Record) Id() string {
	return a.Asset.Id
}

type InMemoryRepository struct {
	repository in_memory.Repository
}

func (r *InMemoryRepository) Create(asset assetEntity.Asset) error {
	record := Record{asset}
	return r.repository.Create(record)
}

func (r *InMemoryRepository) CreateBulk(assets []assetEntity.Asset) (int, error) {
	var records []in_memory.Record
	for _, payment := range assets {
		records = append(records, Record{Asset: payment})
	}
	return r.repository.CreateBulk(records)
}

func NewRepository() *InMemoryRepository {
	return &InMemoryRepository{in_memory.NewRepository()}
}
