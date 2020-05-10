package payment_filters

import (
	"investor/entities/payment"
	"investor/interactors/ports"
	"time"
)

type AssetNameFilter struct {
	paymentFinder ports.PaymentFinderByAssetName
}

func NewAssetNameFilter(finder ports.PaymentFinderByAssetName) AssetNameFilter {
	return AssetNameFilter{paymentFinder: finder}
}

type AssetNameFilterRequest struct {
	TimeFrom  time.Time
	TimeUntil time.Time
	AssetName string
}

type AssetNameFilterResponse struct {
	Payments []payment.Payment
}

func (f AssetNameFilter) Filter(model AssetNameFilterRequest) (AssetNameFilterResponse, error) {
	period := payment.NewDurationPeriod(model.TimeFrom, model.TimeUntil)
	payments, err := f.paymentFinder.FindByAssetName(model.AssetName, period)
	if err != nil {
		return AssetNameFilterResponse{}, err
	}
	return AssetNameFilterResponse{Payments: payments}, nil
}
