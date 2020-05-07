package in_memory

import (
	"errors"
	"investor/entities/asset"
	"investor/entities/payment"
	"reflect"
	"testing"
)

func createRepository() *PaymentRepository {
	finderMock := AssetFinderMock{findFunc: func(assetId string) (a asset.Asset, err error) {
		return asset.Asset{Id: "1", Category: asset.PreciousMetal, Name: "test"}, nil
	}}
	repository := NewPaymentRepository(finderMock)
	return repository
}

func createRepositoryWithBrokenAssetFinder() *PaymentRepository {
	finderMock := AssetFinderMock{findFunc: func(assetId string) (a asset.Asset, err error) {
		return a, AssetDoesntExistsError{AssetId: "mocked_id"}
	}}
	repository := NewPaymentRepository(finderMock)
	return repository
}

func TestPaymentRepository_Create(t *testing.T) {
	repository := createRepository()
	p := payment.CreatePayment("1", 2020)
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
	p1 := payment.CreatePayment("1", 2020)
	p2 := payment.CreatePayment("2", 2020)
	repository := createRepository()

	err := repository.CreateBulk([]payment.Payment{p1, p2})
	if err != nil {
		t.Errorf("Unpected error")
		if len(repository.records) != 2 {
			t.Errorf("2 payments expected to be created")
		}
	}

	repository = createRepository()

	expectedErr := PaymentBulkCreateError{
		FailedIndex: 1, Quantity: 2, Err: PaymentAlreadyExistsError{PaymentId: "1"},
	}
	err = repository.CreateBulk([]payment.Payment{p1, p1})
	if err != expectedErr {
		t.Errorf("Payment alread exists error expected")
	}
	if len(repository.records) != 1 {
		t.Errorf("One payment expected to be created before error")
	}
}

func TestPaymentRepository_ListAll(t *testing.T) {
	records := map[string]PaymentRecord{
		"4": CreatePaymentRecord("4", 2017),
		"3": CreatePaymentRecord("3", 2018),
		"1": CreatePaymentRecord("1", 2020),
		"2": CreatePaymentRecord("2", 2019),
	}
	expectedIds := []string{"1", "2", "3", "4"}

	repository := createRepository()
	repository.records = records

	payments, err := repository.ListAll()
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if len(payments) != 4 {
			t.Errorf("Four payments expected")
		}

		var paymentsIds []string
		for _, p := range payments {
			paymentsIds = append(paymentsIds, p.Id)
		}

		if !reflect.DeepEqual(paymentsIds, expectedIds) {
			t.Errorf("Payments should sorted by date, from the latest to earlier")
		}
	}

	// test when Asset doesnt exists
	repository = createRepositoryWithBrokenAssetFinder()
	repository.records = records

	_, err = repository.ListAll()
	expectedErr := AssetDoesntExistsError{AssetId: "mocked_id"}
	if errors.Is(errors.Unwrap(errors.Unwrap(err)), expectedErr) {
	} else {
		t.Errorf("asset with mocked id doesnt exists error expected")
	}

}
