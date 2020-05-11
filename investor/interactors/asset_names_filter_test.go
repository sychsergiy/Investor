package interactors

import (
	"fmt"
	"investor/entities/payment"
	"reflect"
	"testing"
)

func TestPaymentAssetNamesFilter_Filter(t *testing.T) {
	payments := []payment.Payment{
		payment.CreatePayment("1", 2018),
		payment.CreatePayment("2", 2020),
	}
	mock := PaymentFinderByAssetNamesMock{
		ReturnPayments: payments,
	}
	interactor := NewPaymentAssetNamesFilter(mock)

	req := AssetNamesFilterRequest{
		Periods:    []payment.Period{payment.NewYearPeriod(2020)},
		AssetNames: []string{"test"},
	}
	resp, err := interactor.Filter(req)
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if !reflect.DeepEqual(resp.Payments, payments) {
			t.Errorf("Unexpected payments value")
		}
	}

	mock = PaymentFinderByAssetNamesMock{
		ReturnErr: fmt.Errorf("mocked error"),
	}
	interactor = NewPaymentAssetNamesFilter(mock)
	_, err = interactor.Filter(req)
	if err == nil {
		t.Errorf("error expected")
	}
}
