package ports

import (
	"investor/entities"
	"investor/entities/asset"
	"testing"
	"time"
)

func TestInMemoryStorage_Create(t *testing.T) {
	storage := NewInMemoryStorage()
	testAsset := asset.Asset{Category: asset.CryptoCurrency, Name: "test"}
	creationTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)
	payment := entities.NewReturnPayment("1", 0, 0, testAsset, creationTime)

	// save first payment, no errors expected
	err := storage.Create(payment)
	if err != nil {
		t.Errorf("Unepxected error during payment creation: %s", err)
	}

	// try to save payment with the same id
	err = storage.Create(entities.NewReturnPayment("1", 0, 0, testAsset, creationTime))
	expectedErr := PaymentAlreadyExitsError{"1"}
	if err != expectedErr {
		t.Error("Payment with id already exists error expected")
	}
}
