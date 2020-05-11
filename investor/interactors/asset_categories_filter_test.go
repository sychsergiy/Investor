package interactors

import (
	"errors"
	"investor/entities/asset"
	"investor/entities/payment"
	"reflect"
	"testing"
)

func TestPaymentAssetCategoriesFilter_Filter(t *testing.T) {
	payments := []payment.Payment{payment.CreatePayment("1", 2020)}
	filter := NewPaymentAssetCategoriesFilter(
		PaymentFinderByAssetCategoriesMock{
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

	filter2 := NewPaymentAssetCategoriesFilter(
		PaymentFinderByAssetCategoriesMock{
			ReturnPayments: nil, ReturnErr: asset.NotFoundError{AssetID: "test"},
		},
	)

	_, err = filter2.Filter(AssetCategoriesFilterRequest{})
	if !errors.Is(err, asset.NotFoundError{AssetID: "test"}) {
		t.Errorf("Mocked error expted")
	}
}
