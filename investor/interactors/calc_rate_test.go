package interactors

import (
	"errors"
	"investor/entities/payment"
	"investor/entities/profit"
	"reflect"
	"testing"
)

func TestCalcRateFromProfit_CalcRate(t *testing.T) {
	type unit struct {
		payments         []payment.Payment
		desirableProfit  profit.Profit
		expectedResponse CalcRateFromProfitResponse
		expectedErr      error
	}

	assetName := "test"

	units := []unit{
		{payments: []payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Return, 100, 1),
		},
			desirableProfit: profit.NewFromCoefficient(2),
			expectedErr:     nil,
			expectedResponse: CalcRateFromProfitResponse{
				AssetName:     assetName,
				PaymentsCount: 3,
				AssetRate:     300,
			},
		},
		{payments: []payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 200, 1),
			payment.CreatePaymentWithAmount(payment.Return, 25, 0.75),
			payment.CreatePaymentWithAmount(payment.Return, 25, 0.15),
		},
			desirableProfit: profit.NewFromCoefficient(2),
			expectedErr:     nil,
			expectedResponse: CalcRateFromProfitResponse{
				AssetName:     assetName,
				PaymentsCount: 3,
				AssetRate:     3500,
			},
		},
		{payments: []payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 200, 1),
			payment.CreatePaymentWithAmount(payment.Return, 100, 1),
		},
			desirableProfit: profit.NewFromCoefficient(2),
			expectedErr:     profit.LessThanZeroAssetRestError{},
		},
	}

	for i, u := range units {
		interactor := NewCalcRateFromProfit(PaymentFinderByAssetNamesMock{ReturnPayments: u.payments})
		resp, err := interactor.Calc(
			CalcRateFromProfitRequest{
				AssetName:       assetName,
				Periods:         []payment.Period{},
				DesirableProfit: u.desirableProfit,
			},
		)
		if !errors.Is(err, u.expectedErr) {
			t.Errorf("Unit %d failed. Expected err: %+v. But got: %+v", i, u.expectedErr, err)
		}
		if !reflect.DeepEqual(resp, u.expectedResponse) {
			t.Errorf("Unit %d failed. Unexpected resp value", i)
		}
	}

}
