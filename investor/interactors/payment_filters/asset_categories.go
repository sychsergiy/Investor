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
	Periods         []Period
	PaymentTypes    []payment.Type
	AssetCategories []asset.Category
}

type AssetCategoriesFilterResponse struct {
	Payments []payment.Payment
}

func (f AssetCategoriesFilter) Filter(model AssetCategoriesFilterRequest) (AssetCategoriesFilterResponse, error) {
	var periods []payment.Period
	for _, p := range model.Periods {
		periods = append(periods, payment.NewDurationPeriod(p.TimeFrom, p.TimeUntil))
	}
	payments, err := f.paymentFinder.FindByAssetCategories(
		model.AssetCategories, periods, model.PaymentTypes,
	)
	if err != nil {
		return AssetCategoriesFilterResponse{}, err
	}
	return AssetCategoriesFilterResponse{Payments: payments}, nil
}
