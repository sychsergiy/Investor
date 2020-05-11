package in_memory

import (
	"errors"
	"investor/entities/asset"
	"investor/entities/payment"
	"testing"
)

func createPaymentWithType(paymentType payment.Type) payment.Payment {
	assetRecord := CreateAssetRecord("1", "test")
	return payment.NewPlainPayment(
		"1", 0, 0, assetRecord.ToAsset(),
		payment.CreateYearDate(1), paymentType,
	)
}

func createPaymentWithCreationDate(year int) payment.Payment {
	assetRecord := CreateAssetRecord("1", "test")
	return payment.NewPlainPayment(
		"1", 0, 0, assetRecord.ToAsset(),
		payment.CreateYearDate(year), payment.Invest,
	)
}

func TestFilterByTypes(t *testing.T) {
	type unit struct {
		payments          []payment.Payment
		paymentTypes      []payment.Type
		expectedResultLen int
	}
	invests := []payment.Payment{
		createPaymentWithType(payment.Invest),
		createPaymentWithType(payment.Invest),
	}
	returns := []payment.Payment{
		createPaymentWithType(payment.Return),
		createPaymentWithType(payment.Return),
	}
	mixed := []payment.Payment{
		createPaymentWithType(payment.Return),
		createPaymentWithType(payment.Invest),
	}

	units := []unit{
		{invests, []payment.Type{payment.Invest}, 2},
		{invests, []payment.Type{payment.Return}, 0},
		{returns, []payment.Type{payment.Return}, 2},
		{returns, []payment.Type{payment.Invest}, 0},
		{mixed, []payment.Type{payment.Return}, 1},
		{mixed, []payment.Type{payment.Invest}, 1},
		// return all payments when both payment Types passed
		{invests, []payment.Type{payment.Invest, payment.Return}, 2},
		{mixed, []payment.Type{payment.Invest, payment.Return}, 2},
		{returns, []payment.Type{payment.Invest, payment.Return}, 2},
		// return all payments on empty payments List
		{invests, []payment.Type{}, 2},
		{mixed, []payment.Type{}, 2},
		{returns, []payment.Type{}, 2},
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
		payments          []payment.Payment
		periods           []payment.Period
		expectedResultLen int
	}
	p1 := payment.PeriodMock{
		TimeFrom:  payment.CreateYearDate(2010),
		TimeUntil: payment.CreateYearDate(2012),
	}
	p2 := payment.PeriodMock{
		TimeFrom:  payment.CreateYearDate(2014),
		TimeUntil: payment.CreateYearDate(2016),
	}
	p3 := payment.PeriodMock{
		TimeFrom:  payment.CreateYearDate(2019),
		TimeUntil: payment.CreateYearDate(2020),
	}
	payments := []payment.Payment{
		createPaymentWithCreationDate(2009),
		createPaymentWithCreationDate(2011),
		createPaymentWithCreationDate(2013),
		createPaymentWithCreationDate(2015),
		createPaymentWithCreationDate(2017),
	}
	units := []unit{
		{payments, []payment.Period{p1, p2}, 2},
		{payments, []payment.Period{p1}, 1},
		{payments, []payment.Period{p2}, 1},
		{payments, []payment.Period{p3}, 0},
		{payments, []payment.Period{}, 5},
	}
	for _, u := range units {
		res := FilterByPeriods(u.payments, u.periods)
		if len(res) != u.expectedResultLen {
			t.Errorf("Expected res len %d but got %d", u.expectedResultLen, len(res))
		}
	}
}

func createPaymentWithAssetCategory(category asset.Category) payment.Payment {
	return payment.NewPlainPayment(
		"1", 0, 0,
		asset.NewPlainAsset("1", category, "test"),
		payment.CreateYearDate(1), payment.Invest,
	)
}

func createPaymentWithAssetName(assetName string) payment.Payment {
	return payment.NewPlainPayment(
		"1", 0, 0,
		asset.NewPlainAsset("1", asset.PreciousMetal, assetName),
		payment.CreateYearDate(1), payment.Invest,
	)
}

func TestFilterByAssetCategories(t *testing.T) {
	type unit struct {
		payments          []payment.Payment
		categories        []asset.Category
		expectedResultLen int
	}
	payments := []payment.Payment{
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

	p := NewPaymentProxyMock(
		createPaymentWithAssetCategory(asset.PreciousMetal),
		func() (a asset.Asset, err error) {
			return a, AssetDoesntExistsError{"test"}
		},
	)
	_, err := FilterByAssetCategories([]payment.Payment{p}, []asset.Category{asset.CryptoCurrency})
	if !errors.Is(err, AssetDoesntExistsError{"test"}) {
		t.Errorf("Asset doesnt exist error expected¬")
	}
}

func TestFilterByAssetNames(t *testing.T) {
	type unit struct {
		payments          []payment.Payment
		assetNames        []string
		expectedResultLen int
	}

	asset1 := "asset1"
	asset2 := "asset2"
	payments := []payment.Payment{
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

	p := NewPaymentProxyMock(
		createPaymentWithAssetCategory(asset.PreciousMetal),
		func() (a asset.Asset, err error) {
			return a, AssetDoesntExistsError{"test"}
		},
	)
	_, err := FilterByAssetNames([]payment.Payment{p}, []string{"other"})
	if !errors.Is(err, AssetDoesntExistsError{"test"}) {
		t.Errorf("Asset doesnt exist error expected¬")
	}
}
