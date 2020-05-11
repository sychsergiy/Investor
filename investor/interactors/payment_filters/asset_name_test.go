package payment_filters

import (
	"fmt"
	"investor/entities/payment"
	"investor/interactors"
	"reflect"
	"testing"
)

func TestFilterPayments_Filter(t *testing.T) {
	payments := []payment.Payment{
		payment.CreatePayment("1", 2018),
		payment.CreatePayment("2", 2020),
	}
	mock := interactors.PaymentFinderByAssetNameMock{}
	mock.FindFunc = func(assetName string, period payment.Period) ([]payment.Payment, error) {
		return payments, nil
	}

	interactor := NewAssetNameFilter(mock)

	req := AssetNameFilterRequest{
		TimeFrom:  payment.CreateYearDate(2019),
		TimeUntil: payment.CreateYearDate(2021),
		AssetName: "test",
	}
	resp, err := interactor.Filter(req)
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if !reflect.DeepEqual(resp.Payments, payments) {
			t.Errorf("Unexpected payments value")
		}
	}


	mock.FindFunc = func(assetName string, period payment.Period) ([]payment.Payment, error) {
		return nil, fmt.Errorf("mocked err")
	}
	interactor = NewAssetNameFilter(mock)
	_, err = interactor.Filter(req)
	if err == nil {
		t.Errorf("error expected")
	}
}
