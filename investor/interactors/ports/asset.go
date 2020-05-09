package ports

import "investor/entities/asset"

type AssetCreator interface {
	Create(asset asset.Asset) error
}

type AssetBulkCreator interface {
	CreateBulk(assets []asset.Asset) error
}

type AssetsLister interface {
	ListAll() ([]asset.Asset, error)
}

type AssetRepository interface {
	AssetCreator
	AssetBulkCreator
	AssetsLister
}
