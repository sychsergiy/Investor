package in_memory

import (
	"investor/adapters/repositories"
	assetEntity "investor/entities/asset"
)

type AssetRecord struct {
	Asset assetEntity.Asset
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

func NewInMemoryAssetRepository() *InMemoryAssetRepository {
	return &InMemoryAssetRepository{repositories.NewInMemoryRepository()}
}
