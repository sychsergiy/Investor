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
	FindFunc func(ids []string) ([]payment.Payment, error)
}

type PaymentFinderByAssetNameMock struct {
	FindFunc func(assetName string, period payment.Period) ([]payment.Payment, error)
}

func (m PaymentFinderByAssetNameMock) FindByAssetName(
	assetName string, period payment.Period,
) ([]payment.Payment, error) {
	return m.FindFunc(assetName, period)
}

func (m PaymentFinderByIdsMock) FindByIds(ids []string) ([]payment.Payment, error) {
	return m.FindFunc(ids)
}

type PaymentFinderByAssetCategoriesMock struct {
	ReturnPayments []payment.Payment
	ReturnErr      error
}

func (m PaymentFinderByAssetCategoriesMock) FindByAssetCategories(
	[]asset.Category, []payment.Period, []payment.Type,
) (filtered []payment.Payment, err error) {
	return m.ReturnPayments, m.ReturnErr
}

func (igm IdGeneratorMock) Generate() string {
	return igm.GenerateFunc()
}
