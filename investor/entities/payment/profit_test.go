package payment

import (
	"errors"
	"testing"
)

func TestCalcSumsForPayments(t *testing.T) {
	sums := CalcSumsForPayments([]Payment{
		CreatePaymentWithAmount(Invest, 100, 3),
		CreatePaymentWithAmount(Return, 25, 1),
		CreatePaymentWithAmount(Return, 25, 1),
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
	profit := NewProfitFromPercentage(50)
	if profit.Percentage() != 50 {
		t.Errorf("unexpected percentege value")
	}
	if profit.Coefficient() != 1.5 {
		t.Errorf("unexpected coef value")
	}
}

func TestNewProfitFromCoefficient(t *testing.T) {
	profit := NewProfitFromCoefficient(0.5)
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
	profit, err := CalcProfitForAsset(zeroInvestedSums)
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
	profit, err = CalcProfitForAsset(zeroReturnedSums)
	if !errors.Is(err, expectedErr2) {
		t.Errorf("Zero returned sum error expected")
	}
	if profit != nil {
		t.Errorf("Profit nil value expected")
	}

	profit, err = CalcProfitForAsset(Sums{
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

	profit, err = CalcProfitForAsset(Sums{
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

	profit, err = CalcProfitForAsset(Sums{
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

	profit, err = CalcProfitForAsset(Sums{
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
