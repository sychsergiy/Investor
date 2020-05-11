package payment

import (
	"errors"
	"investor/entities/asset"
	"testing"
)

func TestFilter_ByAssetNames(t *testing.T) {
	type unit struct {
		payments          []Payment
		assetNames        []string
		expectedResultLen int
	}

	asset1 := "asset1"
	asset2 := "asset2"
	payments := []Payment{
		createPaymentWithAssetName(asset1),
		createPaymentWithAssetName(asset2),
	}
	units := []unit{
		{payments, []string{asset1}, 1},
		{payments, []string{asset2}, 1},
		{payments, []string{asset1, asset2}, 2},
		{payments, []string{"other"}, 0},
		{payments, []string{}, 2},
	}
	for _, u := range units {
		filter := NewFilter(u.payments)
		filter, err := filter.ByAssetNames(u.assetNames)
		if err != nil {
			t.Errorf("Unepxected err: %+v", err)
		} else {
			paymentsLen := filter.Payments()
			if len(paymentsLen) != u.expectedResultLen {
				t.Errorf("Expected res len %d but got %d", u.expectedResultLen, paymentsLen)
			}
		}
	}

	p := NewProxyMock(
		createPaymentWithAssetCategory(asset.PreciousMetal),
		func() (a asset.Asset, err error) {
			return a, asset.NotFoundError{AssetID: "test"}
		},
	)
	_, err := FilterByAssetNames([]Payment{p}, []string{"other"})
	if !errors.Is(err, asset.NotFoundError{AssetID: "test"}) {
		t.Errorf("Asset doesnt exist error expected¬")
	}
}

func TestFilter_ByTypes(t *testing.T) {
	type unit struct {
		payments          []Payment
		paymentTypes      []Type
		expectedResultLen int
	}
	invests := []Payment{
		createPaymentWithType(Invest),
		createPaymentWithType(Invest),
	}
	returns := []Payment{
		createPaymentWithType(Return),
		createPaymentWithType(Return),
	}
	mixed := []Payment{
		createPaymentWithType(Return),
		createPaymentWithType(Invest),
	}

	units := []unit{
		{invests, []Type{Invest}, 2},
		{invests, []Type{Return}, 0},
		{returns, []Type{Return}, 2},
		{returns, []Type{Invest}, 0},
		{mixed, []Type{Return}, 1},
		{mixed, []Type{Invest}, 1},
		// return all payments when both payment Types passed
		{invests, []Type{Invest, Return}, 2},
		{mixed, []Type{Invest, Return}, 2},
		{returns, []Type{Invest, Return}, 2},
		// return all payments on empty payments List
		{invests, []Type{}, 2},
		{mixed, []Type{}, 2},
		{returns, []Type{}, 2},
	}

	for _, unit := range units {
		res := NewFilter(unit.payments).ByTypes(unit.paymentTypes).Payments()
		if len(res) != unit.expectedResultLen {
			t.Errorf("Expected result len %d but got %d", unit.expectedResultLen, len(res))
		}
	}
}

func TestFilter_ByPeriods(t *testing.T) {
	type unit struct {
		payments          []Payment
		periods           []Period
		expectedResultLen int
	}
	p1 := PeriodMock{
		TimeFrom:  CreateYearDate(2010),
		TimeUntil: CreateYearDate(2012),
	}
	p2 := PeriodMock{
		TimeFrom:  CreateYearDate(2014),
		TimeUntil: CreateYearDate(2016),
	}
	p3 := PeriodMock{
		TimeFrom:  CreateYearDate(2019),
		TimeUntil: CreateYearDate(2020),
	}
	payments := []Payment{
		createPaymentWithCreationDate(2009),
		createPaymentWithCreationDate(2011),
		createPaymentWithCreationDate(2013),
		createPaymentWithCreationDate(2015),
		createPaymentWithCreationDate(2017),
	}
	units := []unit{
		{payments, []Period{p1, p2}, 2},
		{payments, []Period{p1}, 1},
		{payments, []Period{p2}, 1},
		{payments, []Period{p3}, 0},
		{payments, []Period{}, 5},
	}
	for _, u := range units {
		res := NewFilter(u.payments).ByPeriods(u.periods).Payments()
		if len(res) != u.expectedResultLen {
			t.Errorf("Expected res len %d but got %d", u.expectedResultLen, len(res))
		}
	}
}

func TestFilter_ByAssetCategories(t *testing.T) {
	type unit struct {
		payments          []Payment
		categories        []asset.Category
		expectedResultLen int
	}
	payments := []Payment{
		createPaymentWithAssetCategory(asset.PreciousMetal),
		createPaymentWithAssetCategory(asset.CryptoCurrency),
	}
	units := []unit{
		{payments, []asset.Category{asset.PreciousMetal}, 1},
		{payments, []asset.Category{asset.CryptoCurrency}, 1},
		{payments, []asset.Category{asset.PreciousMetal, asset.CryptoCurrency}, 2},
		{payments, []asset.Category{asset.Stock}, 0},
		{payments, []asset.Category{}, 2},
	}
	for _, u := range units {
		res, err := NewFilter(u.payments).ByAssetCategories(u.categories)
		if err != nil {
			t.Errorf("Unepxected err: %+v", err)
		} else {
			paymentsLen := len(res.Payments())
			if paymentsLen != u.expectedResultLen {
				t.Errorf("Expected res len %d but got %d", u.expectedResultLen, paymentsLen)
			}
		}
	}

	p := NewProxyMock(
		createPaymentWithAssetCategory(asset.PreciousMetal),
		func() (a asset.Asset, err error) {
			return a, asset.NotFoundError{AssetID: "test"}
		},
	)
	_, err := FilterByAssetCategories([]Payment{p}, []asset.Category{asset.CryptoCurrency})
	if !errors.Is(err, asset.NotFoundError{AssetID: "test"}) {
		t.Errorf("Asset doesnt exist error expected¬")
	}
}
