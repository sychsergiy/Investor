package jsonfile

import (
	"fmt"
	"investor/adapters/repositories/memory"
	assetEntity "investor/entities/asset"
)

type AssetRepository struct {
	repository *memory.AssetRepository
	storage    *Storage

	restored bool
}

// todo: change (int, error) to error
func (r *AssetRepository) CreateBulk(assets []assetEntity.Asset) (int, error) {
	err := r.restore()
	if err != nil {
		return 0, err
	}

	n, err := r.repository.CreateBulk(assets)
	if err != nil {
		return n, fmt.Errorf("assets bulk create failed: %w", err)
	}
	err = r.dump()
	return n, err
}

func (r *AssetRepository) Create(a assetEntity.Asset) error {
	err := r.restore()
	if err != nil {
		return err
	}

	err = r.repository.Create(a)
	if err != nil {
		return fmt.Errorf("in memory create asset failed: %w", err)
	}
	return r.dump()
}

func (r *AssetRepository) dump() error {
	assets := r.repository.Records()
	err := r.storage.UpdateAssets(assets)
	if err != nil {
		err = fmt.Errorf("update payments on json storage failed: %w", err)
	}
	return err
}

func (r *AssetRepository) restore() error {
	if r.restored {
		// should be called only once to sync memory storage with file
		return nil
	}
	// read payments from storage file and save in memory
	records, err := r.storage.RetrieveAssets()
	if err != nil {
		return err
	}

	assets := convertRecordsToEntities(records)

	_, err = r.repository.CreateBulk(assets)
	if err != nil {
		err = fmt.Errorf("restore payments failed: %w", err)
	} else {
		r.restored = true
	}
	return err
}

func convertRecordsToEntities(records []memory.AssetRecord) []assetEntity.Asset {
	var assets []assetEntity.Asset
	for _, record := range records {
		assets = append(assets, record.ToAsset())
	}
	return assets
}

func (r *AssetRepository) ListAll() ([]assetEntity.Asset, error) {
	err := r.restore()
	if err != nil {
		return nil, fmt.Errorf("failed to list all assets: %w", err)
	}
	return r.repository.ListAll()
}

func (r *AssetRepository) FindByID(assetID string) (a assetEntity.Asset, err error) {
	err = r.restore()
	if err != nil {
		err = fmt.Errorf("failed to find asset by id: %s due to restore error: %w", assetID, err)
		return
	}
	return r.repository.FindByID(assetID)
}

func NewAssetRepository(s *Storage) *AssetRepository {
	return &AssetRepository{memory.NewAssetRepository(), s, false}
}
