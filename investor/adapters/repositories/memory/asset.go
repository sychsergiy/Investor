package memory

import (
	"fmt"
	"investor/entities/asset"
)

type AssetRecord struct {
	ID       string         `json:"id"`
	Category asset.Category `json:"category"`
	Name     string         `json:"name"`
}

func (ar AssetRecord) ToAsset() asset.Asset {
	return asset.NewPlain(ar.ID, ar.Category, ar.Name)
}

type AssetRecordAlreadyExistsError struct {
	AssetID string
}

func (e AssetRecordAlreadyExistsError) Error() string {
	return fmt.Sprintf("asset with id %s already exists", e.AssetID)
}

func NewAssetRecord(asset asset.Asset) AssetRecord {
	return AssetRecord{asset.ID(), asset.Category(), asset.Name()}
}

type AssetRepository struct {
	records map[string]AssetRecord
}

func (r *AssetRepository) Create(asset asset.Asset) error {
	record := NewAssetRecord(asset)
	_, idExists := r.records[record.ID]

	if idExists {
		return AssetRecordAlreadyExistsError{AssetID: record.ID}
	}
	r.records[record.ID] = record
	return nil
}

// todo: change (int, error) to error
func (r *AssetRepository) CreateBulk(assets []asset.Asset) (int, error) {
	var records []AssetRecord
	for _, a := range assets {
		records = append(records, NewAssetRecord(a))
	}

	var createdCount int
	for createdCount, record := range records {
		_, idExists := r.records[record.ID]
		if idExists {
			return createdCount, AssetRecordAlreadyExistsError{AssetID: record.ID}
		}
		r.records[record.ID] = record
	}
	return createdCount, nil
}

func (r *AssetRepository) ListAll() ([]asset.Asset, error) {
	var assets []asset.Asset
	for _, a := range r.records {
		assets = append(assets, a.ToAsset())
	}
	return assets, nil
}

func (r *AssetRepository) FindByID(assetID string) (a asset.Asset, err error) {
	record, ok := r.records[assetID]
	if !ok {
		err = asset.NotFoundError{AssetID: assetID}
		return
	}
	return record.ToAsset(), nil
}

func (r *AssetRepository) Records() []AssetRecord {
	var records []AssetRecord
	for _, record := range r.records {
		records = append(records, record)
	}
	return records
}

func NewAssetRepository() *AssetRepository {
	return &AssetRepository{make(map[string]AssetRecord)}
}
