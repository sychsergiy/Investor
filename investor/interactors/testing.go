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

type IDGeneratorMock struct {
	GenerateFunc func() string
}

type PaymentFinderByIDsMock struct {
	FindFunc func(ids []string) ([]payment.Payment, error)
}

type PaymentFinderByAssetNamesMock struct {
	ReturnPayments []payment.Payment
	ReturnErr      error
}

func (m PaymentFinderByAssetNamesMock) FindByAssetNames(
	[]string, []payment.Period, []payment.Type,
) ([]payment.Payment, error) {
	return m.ReturnPayments, m.ReturnErr
}

func (m PaymentFinderByIDsMock) FindByIDs(ids []string) ([]payment.Payment, error) {
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

func (igm IDGeneratorMock) Generate() string {
	return igm.GenerateFunc()
}
