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

	p := NewPaymentProxyMock(
		createPaymentWithAssetCategory(asset.PreciousMetal),
		func() (a asset.Asset, err error) {
			return a, asset.AssetDoesntExistsError{AssetId: "test"}
		},
	)
	_, err := FilterByAssetNames([]Payment{p}, []string{"other"})
	if !errors.Is(err, asset.AssetDoesntExistsError{AssetId: "test"}) {
		t.Errorf("Asset doesnt exist error expected¬")
	}
}
