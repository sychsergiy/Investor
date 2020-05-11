package payment_filters

import (
	"investor/entities/asset"
	"investor/entities/payment"
	"investor/interactors/ports"
)

type AssetCategoriesFilter struct {
	paymentFinder ports.PaymentFinderByAssetCategories
}

func NewAssetCategoriesFilter(finder ports.PaymentFinderByAssetCategories) AssetCategoriesFilter {
	return AssetCategoriesFilter{paymentFinder: finder}
}

type AssetCategoriesFilterRequest struct {
	Periods         []payment.Period
	PaymentTypes    []payment.Type
	AssetCategories []asset.Category
}

type AssetCategoriesFilterResponse struct {
	Payments []payment.Payment
}

func (f AssetCategoriesFilter) Filter(model AssetCategoriesFilterRequest) (AssetCategoriesFilterResponse, error) {
	payments, err := f.paymentFinder.FindByAssetCategories(
		model.AssetCategories, model.Periods, model.PaymentTypes,
	)
	if err != nil {
		return AssetCategoriesFilterResponse{}, err
	}
	return AssetCategoriesFilterResponse{Payments: payments}, nil
}
