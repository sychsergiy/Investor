package profit

import (
	"errors"
	"investor/entities/payment"
	"testing"
)

func TestProfitCalculator_CalcForAsset(t *testing.T) {
	type unit struct {
		payments       []payment.Payment
		assetName      string
		expectedProfit AssetProfit
	}
	units := []unit{
		{payments: []payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Return, 200, 1),
		},
			assetName: "test",
			expectedProfit: AssetProfit{
				AssetName:     "test",
				Profit:        NewFromCoefficient(2),
				PaymentsCount: 3,
			},
		},
		{payments: []payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Return, 25, 0.5),
			payment.CreatePaymentWithAmount(payment.Return, 25, 0.5),
		},
			assetName: "test",
			expectedProfit: AssetProfit{
				AssetName:     "test",
				Profit:        NewFromCoefficient(0.5),
				PaymentsCount: 3,
			},
		},
	}

	for _, u := range units {
		profit, err := CalcForAsset(u.payments, u.assetName)
		if err != nil {
			t.Errorf("Unexpected err: %+v", err)
		}
		if profit != u.expectedProfit {
			t.Errorf("Unexpeted profit value")
		}
	}

	type errUnit struct {
		payments    []payment.Payment
		assetName   string
		expectedErr error
	}

	errUnits := []errUnit{
		{[]payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
		},
			"test",
			ZeroAssetReturnedError{},
		},
		{[]payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Return, 25, 1),
			payment.CreatePaymentWithAmount(payment.Return, 25, 1),
		},
			"test",
			ReturnedAssetSumMoreThanInvested{},
		},
	}

	for _, u := range errUnits {
		_, err := CalcForAsset(u.payments, u.assetName)
		if err != u.expectedErr {
			t.Errorf("Expected err: %+v, but got: %+v", err, u.expectedErr)
		}
	}
}

func TestCalcSumsForPayments(t *testing.T) {
	sums := calcSumsForPayments([]payment.Payment{
		payment.CreatePaymentWithAmount(payment.Invest, 100, 3),
		payment.CreatePaymentWithAmount(payment.Return, 25, 1),
		payment.CreatePaymentWithAmount(payment.Return, 25, 1),
	})
	expectedSums := Sums{
		Invested:      100,
		Returned:      50,
		InvestedAsset: 3,
		ReturnedAsset: 2,
	}
	if sums != expectedSums {
		t.Errorf("Unexpected sums value")
	}
}

func TestNewProfitFromPercentage(t *testing.T) {
	profit := NewFromPercentage(50)
	if profit.Percentage() != 50 {
		t.Errorf("unexpected percentege value")
	}
	if profit.Coefficient() != 1.5 {
		t.Errorf("unexpected coef value")
	}
}

func TestNewProfitFromCoefficient(t *testing.T) {
	profit := NewFromCoefficient(0.5)
	if profit.Percentage() != -50 {
		t.Errorf("unexpected percentege value")
	}
	if profit.Coefficient() != 0.5 {
		t.Errorf("unexpected coef value")
	}
}

func TestCalcProfitForAsset(t *testing.T) {
	zeroInvestedSums := Sums{
		Invested:      0,
		Returned:      10,
		InvestedAsset: 10,
		ReturnedAsset: 5,
	}
	expectedErr := ZeroInvestedSumError{}
	profit, err := calcProfitForAsset(zeroInvestedSums)
	if !errors.Is(err, expectedErr) {
		t.Errorf("Zero invested sum error expected")
	}
	if profit != nil {
		t.Errorf("Profit nil value expected")
	}

	zeroReturnedSums := Sums{
		Invested:      10,
		Returned:      5,
		InvestedAsset: 30,
		ReturnedAsset: 0,
	}
	expectedErr2 := ZeroAssetReturnedError{}
	profit, err = calcProfitForAsset(zeroReturnedSums)
	if !errors.Is(err, expectedErr2) {
		t.Errorf("Zero returned sum error expected")
	}
	if profit != nil {
		t.Errorf("Profit nil value expected")
	}

	profit, err = calcProfitForAsset(Sums{
		Invested:      100,
		Returned:      200,
		InvestedAsset: 1,
		ReturnedAsset: 1,
	})
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if profit.Coefficient() != 2 {
			t.Error("2 profit coefficient expected")
		}
		if profit.Percentage() != 100 {
			t.Error("100% profit value expected")
		}
	}

	profit, err = calcProfitForAsset(Sums{
		Invested:      100,
		Returned:      100,
		InvestedAsset: 1,
		ReturnedAsset: 0.5,
	})
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if profit.Coefficient() != 2 {
			t.Error("2 profit coefficient expected")
		}
		if profit.Percentage() != 100 {
			t.Error("100% profit value expected")
		}
	}

	profit, err = calcProfitForAsset(Sums{
		Invested:      400,
		Returned:      100,
		InvestedAsset: 1,
		ReturnedAsset: 0.5,
	})
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if profit.Coefficient() != 0.5 {
			t.Error("0.5 profit coefficient expected")
		}
		if profit.Percentage() != -50 {
			t.Error("-50% profit value expected")
		}
	}

	profit, err = calcProfitForAsset(Sums{
		Invested:      400,
		Returned:      100,
		InvestedAsset: 1,
		ReturnedAsset: 2,
	})
	if !errors.Is(err, ReturnedAssetSumMoreThanInvested{}) {
		t.Errorf("ReturnedAssetSumMoreThanInvested err excpected but got: %+v", err)
	}
	if profit != nil {
		t.Errorf("Profit nil value expected")
	}
}

func TestCalcRateFromDesirableProfit(t *testing.T) {
	type unit struct {
		payments        []payment.Payment
		desirableProfit Profit
		expectedRate    float32
		expectedErr     error
	}
	units := []unit{
		{payments: []payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Invest, 100, 1),
			payment.CreatePaymentWithAmount(payment.Return, 100, 1),
		},
			desirableProfit: NewFromCoefficient(2),
			expectedRate:    300,
		},
		{payments: []payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 200, 1),
			payment.CreatePaymentWithAmount(payment.Return, 25, 0.75),
			payment.CreatePaymentWithAmount(payment.Return, 25, 0.15),
		},
			desirableProfit: NewFromCoefficient(0.5),
			expectedRate:    500,
		},
		{payments: []payment.Payment{
			payment.CreatePaymentWithAmount(payment.Invest, 200, 1),
			payment.CreatePaymentWithAmount(payment.Return, 100, 1),
		},
			desirableProfit: NewFromCoefficient(2),
			expectedErr:     LessThanZeroAssetRestError{},
		},
	}

	for i, u := range units {
		rate, err := CalcRateFromDesirableProfit(u.desirableProfit, u.payments)
		if err != u.expectedErr {
			t.Errorf("Unexpected error on unit %d", i)
		}
		if rate != u.expectedRate {
			t.Errorf("Unexpected rate on unit %d", i)
		}
	}
}
