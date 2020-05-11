package payment_filters

import (
	"investor/entities/payment"
	"investor/interactors/ports"
)

type AssetNamesFilter struct {
	paymentFinder ports.PaymentFinderByAssetNames
}

func NewAssetNamesFilter(finder ports.PaymentFinderByAssetNames) AssetNamesFilter {
	return AssetNamesFilter{paymentFinder: finder}
}

type AssetNameFilterRequest struct {
	Periods      []payment.Period
	PaymentTypes []payment.Type
	AssetNames   []string
}

type AssetNameFilterResponse struct {
	Payments []payment.Payment
}

func (f AssetNamesFilter) Filter(model AssetNameFilterRequest) (AssetNameFilterResponse, error) {
	payments, err := f.paymentFinder.FindByAssetNames(model.AssetNames, model.Periods, model.PaymentTypes)
	if err != nil {
		return AssetNameFilterResponse{}, err
	}
	return AssetNameFilterResponse{Payments: payments}, nil
}
