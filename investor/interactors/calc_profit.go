package interactors

import (
	"investor/entities/payment"
	"investor/entities/profit"
	"investor/interactors/ports"
)

type CalcAssetsProfit struct {
	paymentsFilter ports.PaymentFinderByAssetNames
}

type CalcProfitRequest struct {
	AssetNames []string
	Periods    []payment.Period
}

type CalcAssetsProfitResponse struct {
	Profits []profit.AssetProfit
}

func (cp CalcAssetsProfit) Calc(model CalcProfitRequest) (r CalcAssetsProfitResponse, err error) {
	assetProfits := make([]profit.AssetProfit, 0)

	for _, assetName := range model.AssetNames {
		// todo: create separate method in repository
		payments, err := cp.paymentsFilter.FindByAssetNames(
			[]string{assetName}, model.Periods, []payment.Type{},
		)
		if err != nil {
			return r, err
		}
		calculator := profit.NewProfitCalculator(payments)
		p, err := calculator.CalcForAsset(assetName)
		if err != nil {
			return r, err
		}
		assetProfits = append(assetProfits, p)
	}
	return CalcAssetsProfitResponse{Profits: assetProfits}, nil
}

func NewCalcProfit(paymentsFilter ports.PaymentFinderByAssetNames) CalcAssetsProfit {
	return CalcAssetsProfit{paymentsFilter: paymentsFilter}
}
