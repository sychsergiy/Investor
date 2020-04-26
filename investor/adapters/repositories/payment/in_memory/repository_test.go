package in_memory

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
	expectedErr := PaymentAlreadyExistsError{"1"}
	if err != expectedErr {
		t.Error("Payment with id already exists error expected")
	}
}

func TestInMemoryPaymentRepository_CreateBulk(t *testing.T) {
	creationTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)
	testAsset := asset.Asset{Category: asset.CryptoCurrency, Name: "test"}
	p1 := payment.NewReturn("1", 0, 0, testAsset, creationTime)
	p2 := payment.NewInvestment("2", 0, 0, testAsset, creationTime)
	repository := NewInMemoryPaymentRepository()

	createdQuantity, err := repository.CreateBulk([]payment.Payment{p1, p2})
	if err != nil {
		t.Errorf("Unpected error")
		if createdQuantity != 2 {
			t.Errorf("2 payments expected to be created")
		}
	}

	repository = NewInMemoryPaymentRepository()
	expectedErr := PaymentAlreadyExistsError{"1"}
	createdQuantity, err = repository.CreateBulk([]payment.Payment{p1, p1})
	if err != expectedErr {
		t.Errorf("Payment alread exists error expected")
	}
	if createdQuantity != 1 {
		t.Errorf("One payment expected to be created before error")
	}
}
