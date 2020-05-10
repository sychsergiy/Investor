package payment_filters

import (
	"investor/entities/payment"
	"investor/interactors/ports"
	"time"
)

type AssetNameFilter struct {
	paymentFinder ports.PaymentFinderByAssetName
}

func NewFilterPayments(finder ports.PaymentFinderByAssetName) AssetNameFilter {
	return AssetNameFilter{paymentFinder: finder}
}

type FilterPaymentsRequest struct {
	TimeFrom  time.Time
	TimeUntil time.Time
	AssetName string
}

type FilterPaymentsResponse struct {
	Payments []payment.Payment
}

func (f AssetNameFilter) Filter(model FilterPaymentsRequest) (FilterPaymentsResponse, error) {
	period := payment.NewDurationPeriod(model.TimeFrom, model.TimeUntil)
	payments, err := f.paymentFinder.FindByAssetName(model.AssetName, period)
	if err != nil {
		return FilterPaymentsResponse{}, err
	}
	return FilterPaymentsResponse{Payments: payments}, nil
}
