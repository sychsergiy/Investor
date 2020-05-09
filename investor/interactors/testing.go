package interactors

import "investor/entities/asset"

type AssetCreatorMock struct {
	CreateFunc func(asset asset.Asset) error
}

func (acm AssetCreatorMock) Create(asset asset.Asset) error {
	return acm.CreateFunc(asset)
}

type IdGeneratorMock struct {
	GenerateFunc func() string
}

func (igm IdGeneratorMock) Generate() string {
	return igm.GenerateFunc()
}
