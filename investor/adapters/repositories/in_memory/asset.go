package in_memory

import (
	assetEntity "investor/entities/asset"
)

type AssetRecord struct {
	assetEntity.Asset
}

func (a AssetRecord) Id() string {
	return a.Asset.Id
}

type AssetRepository struct {
	repository Repository
}

func (r *AssetRepository) Create(asset assetEntity.Asset) error {
	record := AssetRecord{asset}
	return r.repository.Create(record)
}

func (r *AssetRepository) CreateBulk(assets []assetEntity.Asset) (int, error) {
	var records []Record
	for _, payment := range assets {
		records = append(records, AssetRecord{Asset: payment})
	}
	return r.repository.CreateBulk(records)
}

func NewAssetRepository() *AssetRepository {
	return &AssetRepository{NewRepository()}
}
