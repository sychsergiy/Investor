package payment

import (
	"errors"
	"investor/adapters/repositories/in_memory"
	"investor/entities/asset"
	"testing"
)

func createPaymentWithType(paymentType Type) Payment {
	assetRecord := in_memory.CreateAssetRecord("1", "test")
	return NewPlainPayment(
		"1", 0, 0, assetRecord.ToAsset(),
		CreateYearDate(1), paymentType,
	)
}

func createPaymentWithCreationDate(year int) Payment {
	assetRecord := in_memory.CreateAssetRecord("1", "test")
	return NewPlainPayment(
		"1", 0, 0, assetRecord.ToAsset(),
		CreateYearDate(year), Invest,
	)
}

func TestFilterByTypes(t *testing.T) {
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
		res := FilterByTypes(unit.payments, unit.paymentTypes)
		if len(res) != unit.expectedResultLen {
			t.Errorf("Expected result len %d but got %d", unit.expectedResultLen, len(res))
		}
	}
}

func TestFilterByPeriods(t *testing.T) {
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
		res := FilterByPeriods(u.payments, u.periods)
		if len(res) != u.expectedResultLen {
			t.Errorf("Expected res len %d but got %d", u.expectedResultLen, len(res))
		}
	}
}

func createPaymentWithAssetCategory(category asset.Category) Payment {
	return NewPlainPayment(
		"1", 0, 0,
		asset.NewPlainAsset("1", category, "test"),
		CreateYearDate(1), Invest,
	)
}

func createPaymentWithAssetName(assetName string) Payment {
	return NewPlainPayment(
		"1", 0, 0,
		asset.NewPlainAsset("1", asset.PreciousMetal, assetName),
		CreateYearDate(1), Invest,
	)
}

func TestFilterByAssetCategories(t *testing.T) {
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
		res, err := FilterByAssetCategories(u.payments, u.categories)
		if err != nil {
			t.Errorf("Unepxected err: %+v", err)
		} else {
			if len(res) != u.expectedResultLen {
				t.Errorf("Expected res len %d but got %d", u.expectedResultLen, len(res))
			}
		}
	}

	p := in_memory.NewPaymentProxyMock(
		createPaymentWithAssetCategory(asset.PreciousMetal),
		func() (a asset.Asset, err error) {
			return a, in_memory.AssetDoesntExistsError{"test"}
		},
	)
	_, err := FilterByAssetCategories([]Payment{p}, []asset.Category{asset.CryptoCurrency})
	if !errors.Is(err, in_memory.AssetDoesntExistsError{"test"}) {
		t.Errorf("Asset doesnt exist error expected¬")
	}
}

func TestFilterByAssetNames(t *testing.T) {
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
		res, err := FilterByAssetNames(u.payments, u.assetNames)
		if err != nil {
			t.Errorf("Unepxected err: %+v", err)
		} else {
			if len(res) != u.expectedResultLen {
				t.Errorf("Expected res len %d but got %d", u.expectedResultLen, len(res))
			}
		}
	}

	p := in_memory.NewPaymentProxyMock(
		createPaymentWithAssetCategory(asset.PreciousMetal),
		func() (a asset.Asset, err error) {
			return a, in_memory.AssetDoesntExistsError{"test"}
		},
	)
	_, err := FilterByAssetNames([]Payment{p}, []string{"other"})
	if !errors.Is(err, in_memory.AssetDoesntExistsError{"test"}) {
		t.Errorf("Asset doesnt exist error expected¬")
	}
}
