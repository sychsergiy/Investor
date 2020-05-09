package interactors

import (
	"errors"
	"investor/entities/asset"
	"investor/entities/payment"
	"testing"
	"time"
)

func createPayment(type_ payment.Type, amount, assetAmount float32) payment.Payment {
	a := asset.NewPlainAsset("gold", asset.PreciousMetal, "gold")
	date := time.Date(2019, 30, 12, 11, 58, 0, 0, time.UTC)
	return payment.NewPlainPayment("test", assetAmount, amount, a, date, type_)
}

func TestCalcProfit_Calc(t *testing.T) {
	calcProfit := NewCalcProfit()
	profit, err := calcProfit.Calc([]payment.Payment{
		createPayment(payment.Invest, 100, 1),
		createPayment(payment.Invest, 100, 1),
		createPayment(payment.Return, 200, 1),
	})
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if profit.Coefficient() != 2 || profit.Percentage() != 100 {
			t.Errorf("Unexpected profit value")
		}
	}

	calcProfit = NewCalcProfit()
	profit, err = calcProfit.Calc([]payment.Payment{
		createPayment(payment.Invest, 100, 1),
		createPayment(payment.Return, 25, 0.5),
		createPayment(payment.Return, 25, 0.5),
	})
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if profit.Coefficient() != 0.5 || profit.Percentage() != -50 {
			t.Errorf("Unexpected profit value")
		}
	}

	profit, err = calcProfit.Calc([]payment.Payment{
		createPayment(payment.Invest, 100, 1),
		createPayment(payment.Invest, 100, 1),
	})
	if !errors.Is(err, payment.ZeroAssetReturnedError{}) {
		t.Errorf("ZeroAssetReturnedError err expected")
	}
	if profit != nil {
		t.Errorf("profit nil value expected")
	}

	profit, err = calcProfit.Calc([]payment.Payment{
		createPayment(payment.Return, 100, 1),
		createPayment(payment.Return, 100, 1),
	})
	if !errors.Is(err, payment.ZeroInvestedSumError{}) {
		t.Errorf("ZeroAssetReturnedError err expected")
	}
	if profit != nil {
		t.Errorf("profit nil value expected")
	}

	calcProfit = NewCalcProfit()
	profit, err = calcProfit.Calc([]payment.Payment{
		createPayment(payment.Invest, 100, 1),
		createPayment(payment.Return, 25, 1),
		createPayment(payment.Return, 25, 1),
	})
	if !errors.Is(err, payment.ReturnedAssetSumMoreThanInvested{}) {
		t.Errorf("Unexpected err: %+v", err)
	}
	if profit != nil {
		t.Errorf("Profit nil value expected")
	}
}

func TestCalcSumsForPayments(t *testing.T) {
	sums := CalcSumsForPayments([]payment.Payment{
		createPayment(payment.Invest, 100, 3),
		createPayment(payment.Return, 25, 1),
		createPayment(payment.Return, 25, 1),
	})
	expectedSums := payment.Sums{
		Invested:      100,
		Returned:      50,
		InvestedAsset: 3,
		ReturnedAsset: 2,
	}
	if sums != expectedSums {
		t.Errorf("Unexpected sums value")
	}
}
