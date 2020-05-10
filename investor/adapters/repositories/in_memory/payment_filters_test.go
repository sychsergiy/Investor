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

func TestFilterPaymentsByType(t *testing.T) {
	type unit struct {
		payments          []payment.Payment
		paymentType       payment.Type
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
		{invests, payment.Invest, 2},
		{invests, payment.Return, 0},
		{returns, payment.Return, 2},
		{returns, payment.Invest, 0},
		{mixed, payment.Return, 1},
		{mixed, payment.Invest, 1},
	}

	for _, unit := range units {
		res := FilterPaymentsByType(unit.payments, unit.paymentType)
		if len(res) != unit.expectedResultLen {
			t.Errorf("Expected result len %d but got %d", unit.expectedResultLen, len(res))
		}
	}
}
