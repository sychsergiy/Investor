package in_memory

import (
	"fmt"
	assetEntity "investor/entities/asset"
)

type InMemoryAssetRepository struct {
	assets map[string]assetEntity.Asset
}

type AssetAlreadyExistsError struct {
	AssetId string
}

func (e AssetAlreadyExistsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.AssetId)
}

func (repository *InMemoryAssetRepository) Create(asset assetEntity.Asset) (err error) {
	_, idExists := repository.assets[asset.Id]
	if idExists {
		err = AssetAlreadyExistsError{AssetId: asset.Id}
	} else {
		repository.assets[asset.Id] = asset
	}
	return
}

func NewInMemoryAssetRepository() *InMemoryAssetRepository {
	return &InMemoryAssetRepository{assets: make(map[string]assetEntity.Asset)}
}
