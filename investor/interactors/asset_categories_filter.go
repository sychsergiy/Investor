package interactors

import (
	"investor/entities/asset"
	"investor/entities/payment"
	"investor/interactors/ports"
)

type PaymentAssetCategoriesFilter struct {
	paymentFinder ports.PaymentFinderByAssetCategories
}

func NewPaymentAssetCategoriesFilter(finder ports.PaymentFinderByAssetCategories) PaymentAssetCategoriesFilter {
	return PaymentAssetCategoriesFilter{paymentFinder: finder}
}

type AssetCategoriesFilterRequest struct {
	Periods         []payment.Period
	PaymentTypes    []payment.Type
	AssetCategories []asset.Category
}

type AssetCategoriesFilterResponse struct {
	Payments []payment.Payment
}

func (f PaymentAssetCategoriesFilter) Filter(model AssetCategoriesFilterRequest) (AssetCategoriesFilterResponse, error) {
	payments, err := f.paymentFinder.FindByAssetCategories(
		model.AssetCategories, model.Periods, model.PaymentTypes,
	)
	if err != nil {
		return AssetCategoriesFilterResponse{}, err
	}
	return AssetCategoriesFilterResponse{Payments: payments}, nil
}
