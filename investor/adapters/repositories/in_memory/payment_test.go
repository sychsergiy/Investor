package in_memory

import (
	"investor/entities/asset"
	"investor/entities/payment"
	"reflect"
	"testing"
	"time"
)

func createPayment(id string, year int) payment.Payment {
	testAsset := asset.Asset{Category: asset.CryptoCurrency, Name: "test"}
	creationTime := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	return payment.NewReturn(id, 0, 0, testAsset, creationTime)
}

func TestPaymentRepository_Create(t *testing.T) {
	repository := NewPaymentRepository()
	p := createPayment("1", 2020)
	// save first payment, no errors expected
	err := repository.Create(p)
	if err != nil {
		t.Errorf("Unepxected error during payment creation: %s", err)
	}

	// try to save payment with the same id
	err = repository.Create(p)
	expectedErr := PaymentAlreadyExistsError{PaymentId: p.Id}
	if err != expectedErr {
		t.Error("Payment with id already exists error expected")
	}
}

func TestPaymentRepository_CreateBulk(t *testing.T) {
	p1 := createPayment("1", 2020)
	p2 := createPayment("2", 2020)
	repository := NewPaymentRepository()

	createdQuantity, err := repository.CreateBulk([]payment.Payment{p1, p2})
	if err != nil {
		t.Errorf("Unpected error")
		if createdQuantity != 2 {
			t.Errorf("2 payments expected to be created")
		}
	}

	repository = NewPaymentRepository()
	expectedErr := RecordAlreadyExistsError{RecordId: "1"}
	createdQuantity, err = repository.CreateBulk([]payment.Payment{p1, p1})
	if err != expectedErr {
		t.Errorf("Payment alread exists error expected")
	}
	if createdQuantity != 1 {
		t.Errorf("One payment expected to be created before error")
	}
}

func TestPaymentRepository_ListAll(t *testing.T) {
	records := map[string]payment.Payment{
		"4": createPayment("4", 2017),
		"3": createPayment("3", 2018),
		"1": createPayment("1", 2020),
		"2": createPayment("2", 2019),
	}
	repository := PaymentRepository{records}
	payments := repository.ListAll()
	if len(payments) != 4 {
		t.Errorf("Four payments expected")
	}

	var paymentsIds []string
	for _, p := range payments {
		paymentsIds = append(paymentsIds, p.Id)
	}

	expectedIds := []string{"1", "2", "3", "4"}
	if !reflect.DeepEqual(paymentsIds, expectedIds) {
		t.Errorf("Payments should sorted by date, from the latest to earlier")
	}
}
