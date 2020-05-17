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
	assetName       string
	periods         []payment.Period
	desirableProfit profit.Profit
}

type CalcRateFromProfitResponse struct {
	assetName     string
	paymentsCount int
	assetRate     float32
}

func (c CalcRateFromProfit) Calc(model CalcRateFromProfitRequest) (
	CalcRateFromProfitResponse, error,
) {
	payments, err := c.paymentsFilter.FindByAssetNames(
		[]string{model.assetName},
		model.periods,
		[]payment.Type{},
	)
	if err != nil {
		return CalcRateFromProfitResponse{}, err
	}
	rate, err := profit.CalcRateFromDesirableProfit(model.desirableProfit, payments)
	if err != nil {
		return CalcRateFromProfitResponse{}, err
	}
	return CalcRateFromProfitResponse{
		assetName:     model.assetName,
		paymentsCount: len(payments),
		assetRate:     rate,
	}, nil
}

func NewCalcRateFromProfit(
	paymentsFilter ports.PaymentFinderByAssetNames,
) CalcRateFromProfit {
	return CalcRateFromProfit{paymentsFilter: paymentsFilter}
}
