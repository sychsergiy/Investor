package asset

import (
	"investor/adapters/repositories/in_memory"
	"investor/adapters/repositories/in_memory/asset"
	"investor/adapters/repositories/jsonfile"
	assetEntity "investor/entities/asset"
)

type JsonFileRepository struct {
	repository jsonfile.Repository
}

func (r JsonFileRepository) CreateBulk(assets []assetEntity.Asset) (int, error) {
	var records []in_memory.Record
	for _, a := range assets {
		records = append(records, asset.Record{Asset: a})
	}
	return r.repository.CreateBulk(records)
}

func (r JsonFileRepository) Create(a assetEntity.Asset) error {
	return r.repository.Create(asset.Record{Asset: a})
}

func NewRepository(repository jsonfile.Repository) JsonFileRepository {
	return JsonFileRepository{repository}
}
