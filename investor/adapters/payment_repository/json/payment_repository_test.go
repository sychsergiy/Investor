package json

import (
	"investor/entities/asset"
	"investor/entities/payment"
	"testing"
	"time"
)

func TestInMemoryPaymentRepository_Create(t *testing.T) {
	repository := NewInMemoryPaymentRepository()
	testAsset := asset.Asset{Category: asset.CryptoCurrency, Name: "test"}
	creationTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)
	p := payment.NewReturn("1", 0, 0, testAsset, creationTime)

	// save first payment, no errors expected
	err := repository.Create(p)
	if err != nil {
		t.Errorf("Unepxected error during payment creation: %s", err)
	}

	// try to save payment with the same id
	err = repository.Create(p)
	expectedErr := PaymentAlreadyExitsError{"1"}
	if err != expectedErr {
		t.Error("Payment with id already exists error expected")
	}
}
