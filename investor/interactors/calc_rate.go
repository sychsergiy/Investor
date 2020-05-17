package interactors

import (
	"investor/entities/payment"
	"investor/entities/profit"
	"investor/interactors/ports"
)

type CalcRateFromProfit struct {
	paymentsFilter ports.PaymentFinderByAssetNames
}

type CalcRateFromProfitRequest struct {
	AssetName       string
	Periods         []payment.Period
	DesirableProfit profit.Profit
}

type CalcRateFromProfitResponse struct {
	AssetName     string
	PaymentsCount int
	AssetRate     float32
}

func (c CalcRateFromProfit) Calc(model CalcRateFromProfitRequest) (
	CalcRateFromProfitResponse, error,
) {
	payments, err := c.paymentsFilter.FindByAssetNames(
		[]string{model.AssetName},
		model.Periods,
		[]payment.Type{},
	)
	if err != nil {
		return CalcRateFromProfitResponse{}, err
	}
	rate, err := profit.CalcRateFromDesirableProfit(model.DesirableProfit, payments)
	if err != nil {
		return CalcRateFromProfitResponse{}, err
	}
	return CalcRateFromProfitResponse{
		AssetName:     model.AssetName,
		PaymentsCount: len(payments),
		AssetRate:     rate,
	}, nil
}

func NewCalcRateFromProfit(
	paymentsFilter ports.PaymentFinderByAssetNames,
) CalcRateFromProfit {
	return CalcRateFromProfit{paymentsFilter: paymentsFilter}
}
