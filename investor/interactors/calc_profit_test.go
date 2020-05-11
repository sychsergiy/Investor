package interactors

import (
	"errors"
	"investor/entities/payment"
	"reflect"
	"testing"
)

func TestCalcProfit_Calc(t *testing.T) {
	type unit struct {
		payments     []payment.Payment
		assetNames   []string
		expectedResp CalcAssetsProfitResponse
		expectedErr  error
	}
	units := []unit{
		{
			payments: []payment.Payment{
				payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
				payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
				payment.CreatePaymentWithAmount(payment.Return, 200, 1),
			},
			assetNames: []string{"test"},
			expectedResp: CalcAssetsProfitResponse{
				Profits: []payment.AssetProfit{
					{
						AssetName:     "test",
						Profit:        payment.NewProfitFromCoefficient(2),
						PaymentsCount: 3,
					},
				},
			},
			expectedErr: nil,
		},
		{
			payments: []payment.Payment{
				payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
				payment.CreatePaymentWithAmount(payment.Return, 25, 0.5),
				payment.CreatePaymentWithAmount(payment.Return, 25, 0.5),
			},
			assetNames: []string{"test"},
			expectedResp: CalcAssetsProfitResponse{
				[]payment.AssetProfit{
					{
						AssetName:     "test",
						Profit:        payment.NewProfitFromCoefficient(0.5),
						PaymentsCount: 3,
					},
				},
			},
			expectedErr: nil,
		},
		{
			payments: []payment.Payment{
				payment.CreatePaymentWithAmount(payment.Return, 100, 1),
				payment.CreatePaymentWithAmount(payment.Return, 100, 1),
			},
			assetNames:   []string{"test"},
			expectedResp: CalcAssetsProfitResponse{},
			expectedErr:  payment.ZeroInvestedSumError{},
		},
		{
			payments: []payment.Payment{
				payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
				payment.CreatePaymentWithAmount(payment.Return, 25, 1),
				payment.CreatePaymentWithAmount(payment.Return, 25, 1),
			},
			assetNames:   []string{"test"},
			expectedResp: CalcAssetsProfitResponse{},
			expectedErr:  payment.ReturnedAssetSumMoreThanInvested{},
		},
	}

	for i, u := range units {
		interactor := NewCalcProfit(PaymentFinderByAssetNamesMock{ReturnPayments: u.payments})
		resp, err := interactor.Calc(CalcProfitRequest{AssetNames: u.assetNames})
		if !errors.Is(err, u.expectedErr) {
			t.Errorf("Unit %d failed. Expected err: %+v. But got: %+v", i, u.expectedErr, err)
		}
		if !reflect.DeepEqual(resp, u.expectedResp) {
			t.Errorf("Unit %d failed. Unexpected resp value", i)
		}
	}
}
