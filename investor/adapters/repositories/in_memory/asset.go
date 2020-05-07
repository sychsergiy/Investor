package in_memory

import (
	"fmt"
	assetEntity "investor/entities/asset"
)

type AssetRecord struct {
	Id       string `json:"id"`
	Category int    `json:"category"`
	Name     string `json:"name"`
}

func (ar AssetRecord) ToAsset() assetEntity.Asset {
	return assetEntity.Asset{Id: ar.Id, Category: assetEntity.Category(ar.Category), Name: ar.Name}
}

type AssetRecordAlreadyExistsError struct {
	RecordId string
}

func (e AssetRecordAlreadyExistsError) Error() string {
	return fmt.Sprintf("asset with id %s already exists", e.RecordId)

}

func NewAssetRecord(asset assetEntity.Asset) AssetRecord {
	return AssetRecord{asset.Id, int(asset.Category), asset.Name}
}

type AssetRepository struct {
	records map[string]AssetRecord
}

func (r *AssetRepository) Create(asset assetEntity.Asset) error {
	record := NewAssetRecord(asset)
	_, idExists := r.records[record.Id]

	if idExists {
		return AssetRecordAlreadyExistsError{RecordId: record.Id}
	} else {
		r.records[record.Id] = record
		return nil
	}
}

func (r *AssetRepository) CreateBulk(assets []assetEntity.Asset) (int, error) {
	var records []AssetRecord
	for _, a := range assets {
		records = append(records, NewAssetRecord(a))
	}

	var createdCount int
	for createdCount, record := range records {
		_, idExists := r.records[record.Id]
		if idExists {
			return createdCount, AssetRecordAlreadyExistsError{RecordId: record.Id}
		} else {
			r.records[record.Id] = record
		}
	}
	return createdCount, nil
}

func (r *AssetRepository) ListAll() ([]assetEntity.Asset, error) {
	var payments []assetEntity.Asset
	for _, a := range r.records {
		payments = append(payments, a.ToAsset())
	}
	return payments, nil
}

func NewAssetRepository() *AssetRepository {
	return &AssetRepository{make(map[string]AssetRecord)}
}
