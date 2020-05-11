package in_memory

import (
	"investor/entities/payment"
	"testing"
)

func createPaymentWithType(paymentType payment.Type) payment.Payment {
	assetRecord := CreateAssetRecord("1", "test")
	return payment.NewPlainPayment(
		"1", 0, 0, assetRecord.ToAsset(),
		payment.CreateYearDate(2020), paymentType,
	)
}

func createPaymentWithCreationDate(year int) payment.Payment {
	assetRecord := CreateAssetRecord("1", "test")
	return payment.NewPlainPayment(
		"1", 0, 0, assetRecord.ToAsset(),
		payment.CreateYearDate(year), payment.Invest,
	)
}

func TestFilterByType(t *testing.T) {
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
		res := FilterByType(unit.payments, unit.paymentTypes)
		if len(res) != unit.expectedResultLen {
			t.Errorf("Expected result len %d but got %d", unit.expectedResultLen, len(res))
		}
	}
}

func TestFilterByPeriod(t *testing.T) {
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
		{payments, []payment.Period{}, 5},
	}
	for _, u := range units {
		res := FilterByPeriod(u.payments, u.periods)
		if len(res) != u.expectedResultLen {
			t.Errorf("Expected res len %d but got %d", u.expectedResultLen, len(res))
		}
	}
}
