package jsonfile

import (
	"investor/adapters/repositories/in_memory"
	assetEntity "investor/entities/asset"
)

type AssetRepository struct {
	repository Repository
}

func (r AssetRepository) CreateBulk(assets []assetEntity.Asset) (int, error) {
	var records []in_memory.Record
	for _, a := range assets {
		records = append(records, in_memory.AssetRecord{Asset: a})
	}
	return r.repository.CreateBulk(records)
}

func (r AssetRepository) Create(a assetEntity.Asset) error {
	return r.repository.Create(in_memory.AssetRecord{Asset: a})
}

func NewAssetRepository(repository Repository) AssetRepository {
	return AssetRepository{repository}
}
