package payment_filters

import (
	"errors"
	"investor/adapters/repositories/in_memory"
	"investor/entities/payment"
	"investor/interactors"
	"reflect"
	"testing"
)

func TestAssetCategoriesFilter_Filter(t *testing.T) {
	payments := []payment.Payment{payment.CreatePayment("1", 2020)}
	filter := NewAssetCategoriesFilter(
		interactors.PaymentFinderByAssetCategoriesMock{
			ReturnPayments: payments, ReturnErr: nil,
		},
	)
	resp, err := filter.Filter(AssetCategoriesFilterRequest{})
	if err != nil {
		t.Errorf("Unexpted err: %+v", err)
	}
	if !reflect.DeepEqual(resp.Payments, payments) {
		t.Errorf("Unexpted returne payments value")
	}

	filter2 := NewAssetCategoriesFilter(
		interactors.PaymentFinderByAssetCategoriesMock{
			ReturnPayments: nil, ReturnErr: in_memory.AssetDoesntExistsError{AssetId: "test"},
		},
	)

	_, err = filter2.Filter(AssetCategoriesFilterRequest{})
	if !errors.Is(err, in_memory.AssetDoesntExistsError{AssetId: "test"}) {
		t.Errorf("Mocked error expted")
	}
}
