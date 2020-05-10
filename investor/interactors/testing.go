package interactors

import (
	"investor/entities/asset"
	"investor/entities/payment"
)

type AssetCreatorMock struct {
	CreateFunc func(asset asset.Asset) error
}

func (acm AssetCreatorMock) Create(asset asset.Asset) error {
	return acm.CreateFunc(asset)
}

type IdGeneratorMock struct {
	GenerateFunc func() string
}

type PaymentFinderByIdsMock struct {
	FindByIdsFunc func(ids []string) ([]payment.Payment, error)
}

func (m PaymentFinderByIdsMock) FindByIds(ids []string) ([]payment.Payment, error) {
	return m.FindByIdsFunc(ids)
}

func (igm IdGeneratorMock) Generate() string {
	return igm.GenerateFunc()
}
