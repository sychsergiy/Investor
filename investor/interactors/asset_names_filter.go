package interactors

import (
	"investor/entities/payment"
	"investor/interactors/ports"
)

type PaymentAssetNamesFilter struct {
	paymentFinder ports.PaymentFinderByAssetNames
}

func NewPaymentAssetNamesFilter(finder ports.PaymentFinderByAssetNames) PaymentAssetNamesFilter {
	return PaymentAssetNamesFilter{paymentFinder: finder}
}

type AssetNamesFilterRequest struct {
	Periods      []payment.Period
	PaymentTypes []payment.Type
	AssetNames   []string
}

type AssetNameFilterResponse struct {
	Payments []payment.Payment
}

func (f PaymentAssetNamesFilter) Filter(model AssetNamesFilterRequest) (AssetNameFilterResponse, error) {
	payments, err := f.paymentFinder.FindByAssetNames(model.AssetNames, model.Periods, model.PaymentTypes)
	if err != nil {
		return AssetNameFilterResponse{}, err
	}
	return AssetNameFilterResponse{Payments: payments}, nil
}
