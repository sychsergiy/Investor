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

type AssetRecordAlreadyExistsError struct {
	RecordId string
}

func (e AssetRecordAlreadyExistsError) Error() string {
	return fmt.Sprintf("asset with id %s already exists", e.RecordId)

}

func newAssetRecord(asset assetEntity.Asset) AssetRecord {
	return AssetRecord{asset.Id, int(asset.Category), asset.Name}
}

type AssetRepository struct {
	records map[string]AssetRecord
}

func (r *AssetRepository) Create(asset assetEntity.Asset) error {
	record := newAssetRecord(asset)
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
		records = append(records, newAssetRecord(a))
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

func NewAssetRepository() *AssetRepository {
	return &AssetRepository{make(map[string]AssetRecord)}
}
