package in_memory

import (
	"investor/adapters/repositories"
	assetEntity "investor/entities/asset"
)

type AssetRecord struct {
	assetEntity.Asset
}

func (a AssetRecord) Id() string {
	return a.Asset.Id
}

type InMemoryAssetRepository struct {
	repository repositories.InMemoryRepository
}

func (r *InMemoryAssetRepository) Create(asset assetEntity.Asset) error {
	record := AssetRecord{asset}
	return r.repository.Create(record)
}

func (r *InMemoryAssetRepository) CreateBulk(assets []assetEntity.Asset) (int, error) {
	var records []repositories.Record
	for _, payment := range assets {
		records = append(records, AssetRecord{Asset: payment})
	}
	return r.repository.CreateBulk(records)
}

func NewInMemoryAssetRepository() *InMemoryAssetRepository {
	return &InMemoryAssetRepository{repositories.NewInMemoryRepository()}
}
