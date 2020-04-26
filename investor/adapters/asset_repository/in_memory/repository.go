package in_memory

import (
	"fmt"
	assetEntity "investor/entities/asset"
)

type InMemoryAssetRepository struct {
	payments map[string]assetEntity.Asset
}

type AssetAlreadyExitsError struct {
	AssetId string
}

func (e AssetAlreadyExitsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.AssetId)
}

func (storage *InMemoryAssetRepository) Create(asset assetEntity.Asset) (err error) {
	_, idExists := storage.payments[asset.Id]
	if idExists {
		err = AssetAlreadyExitsError{AssetId: asset.Id}
	} else {
		storage.payments[asset.Id] = asset
	}
	return
}

func NewInMemoryAssetRepository() *InMemoryAssetRepository {
	return &InMemoryAssetRepository{payments: make(map[string]assetEntity.Asset)}
}
