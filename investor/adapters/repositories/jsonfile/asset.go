package jsonfile

import (
	"fmt"
	"investor/adapters/repositories/in_memory"
	assetEntity "investor/entities/asset"
)

type AssetRepository struct {
	repository *in_memory.AssetRepository
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
	// todo: use records method here
	assets, err := r.repository.ListAll()
	if err != nil {
		return fmt.Errorf("list assets failed: %w", err)
	}
	err = r.storage.UpdateAssets(assets)
	if err != nil {
		err = fmt.Errorf("update payments on json storage failed: %w", err)
	}
	return err
}

func (r *AssetRepository) restore() error {
	if r.restored {
		// should be called only once to sync in_memory storage with file
		return nil
	}
	// read payments from storage file and save in memory
	assets, err := r.storage.RetrieveAssets()
	if err != nil {
		return err
	}
	_, err = r.repository.CreateBulk(assets)
	if err != nil {
		err = fmt.Errorf("restore payments failed, storage file malformed: %w", err)
	} else {
		r.restored = true
	}
	return err
}

func (r *AssetRepository) ListAll() ([]assetEntity.Asset, error) {
	err := r.restore()
	if err != nil {
		return nil, fmt.Errorf("failed to list all assets: %w", err)
	}
	return r.repository.ListAll()
}

func (r *AssetRepository) FindById(assetId string) (a assetEntity.Asset, err error) {
	err = r.restore()
	if err != nil {
		err = fmt.Errorf("failed to find asset by Id: %s due to restore error: %w", assetId, err)
		return
	}
	return r.repository.FindById(assetId)
}

func NewAssetRepository(s *Storage) *AssetRepository {
	return &AssetRepository{in_memory.NewAssetRepository(), s, false}
}
