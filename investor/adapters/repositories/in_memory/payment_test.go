package in_memory

import (
	"errors"
	"investor/entities/asset"
	"investor/entities/payment"
	"reflect"
	"testing"
)

func CreatePaymentWithoutAsset(id string) PaymentProxyMock {
	return NewPaymentProxyMock(
		payment.CreatePayment(id, 2020),
		func() (a asset.Asset, err error) { return a, AssetDoesntExistsError{AssetId: "mocked_id"} },
	)
}
func createRepository() *PaymentRepository {
	finderMock := AssetFinderMock{findFunc: func(assetId string) (a asset.Asset, err error) {
		return asset.NewPlainAsset("1", asset.PreciousMetal, "test"), nil
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
	expectedErr := PaymentAlreadyExistsError{PaymentId: p.Id()}
	if err != expectedErr {
		t.Error("Payment with id already exists error expected")
	}

	repository2 := createRepository()
	p2 := CreatePaymentWithoutAsset("2")
	err2 := repository2.Create(p2)
	expectedErr2 := AssetDoesntExistsError{AssetId: "mocked_id"}
	if !errors.Is(err2, expectedErr2) {
		t.Errorf("Asset with provided doesn't exist error expected")
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

	repository = createRepository()
	err = repository.CreateBulk([]payment.Payment{CreatePaymentWithoutAsset("1"), p2})
	expectedErr = PaymentBulkCreateError{
		FailedIndex: 0, Quantity: 2, Err: AssetDoesntExistsError{AssetId: "mocked_id"},
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("payment bulk create error with asset doesn exists root cause error expected")
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
		paymentsIds := payment.PaymentsToIds(payments)
		if !reflect.DeepEqual(paymentsIds, expectedIds) {
			t.Errorf("Payments should sorted by date, from the latest to earlier")
		}
	}
}

func TestPaymentRepository_FindByIds(t *testing.T) {
	repository := createRepository()
	repository.records = map[string]PaymentRecord{
		"4": CreatePaymentRecord("4", 2017),
		"3": CreatePaymentRecord("3", 2018),
		"1": CreatePaymentRecord("1", 2020),
		"2": CreatePaymentRecord("2", 2019),
	}
	ids := []string{"4", "2", "3"}
	expectedIds := []string{"2", "3", "4"}
	payments, err := repository.FindByIds(ids)
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		paymentsIds := payment.PaymentsToIds(payments)
		if !reflect.DeepEqual(expectedIds, paymentsIds) {
			t.Errorf("Unexpected payments, slice with 3 listered ids epxected")
		}
	}

	payments, err = repository.FindByIds([]string{"not_existent"})
	if !errors.Is(err, PaymentDoesntExistsError{"not_existent"}) {
		t.Errorf("PaymentsDoesntExistsError expected, but got: %+v", err)
	}
	if payments != nil {
		t.Errorf("Payments nil value expected")
	}
}

func createPaymentRecord(paymentType payment.Type, year int, assetId string) PaymentRecord {
	return PaymentRecord{
		Id:             "test",
		AssetAmount:    0,
		AbsoluteAmount: 0,
		AssetId:        assetId,
		Type:           paymentType,
		CreationDate:   payment.CreateYearDate(year),
	}
}

func TestPaymentRepository_FindByAssetCategories(t *testing.T) {
	assetRepo := NewAssetRepository()
	_, err := assetRepo.CreateBulk([]asset.Asset{
		asset.NewPlainAsset("1", asset.PreciousMetal, "gold"),
		asset.NewPlainAsset("2", asset.CryptoCurrency, "BTC"),
	})
	if err != nil {
		t.Errorf("Unexpted err during preparation: %+v", err)
		return
	}

	repository := NewPaymentRepository(assetRepo)
	repository.records = map[string]PaymentRecord{
		"1": createPaymentRecord(payment.Invest, 2000, "1"),
		"2": createPaymentRecord(payment.Return, 2019, "1"),
		"3": createPaymentRecord(payment.Invest, 2019, "2"),
		"4": createPaymentRecord(payment.Invest, 2019, "1"),
	}

	res, err := repository.FindByAssetCategories(
		[]asset.Category{asset.PreciousMetal},
		[]payment.Period{payment.PeriodMock{
			TimeFrom:  payment.CreateYearDate(2018),
			TimeUntil: payment.CreateYearDate(2020),
		}},
		[]payment.Type{payment.Invest},
	)
	if err != nil {
		t.Errorf("Unexpeted err: %+v", err)
	}
	if len(res) != 1 {
		t.Errorf("Unepxted result, 1 item expected but got %d", len(res))
	}

	res, err = repository.FindByAssetCategories(
		[]asset.Category{},
		[]payment.Period{},
		[]payment.Type{},
	)
	if err != nil {
		t.Errorf("Unexpeted err: %+v", err)
	}
	if len(res) != 4 {
		t.Errorf("All 4 items exected to be returned when no filters passed")
	}
}

func TestPaymentRepository_FindByAssetNames(t *testing.T) {
	assetRepo := NewAssetRepository()
	gold := "gold"
	bitcoin := "BTC"
	_, err := assetRepo.CreateBulk([]asset.Asset{
		asset.NewPlainAsset("1", asset.PreciousMetal, gold),
		asset.NewPlainAsset("2", asset.CryptoCurrency, bitcoin),
	})
	if err != nil {
		t.Errorf("Unexpted err during preparation: %+v", err)
		return
	}

	repository := NewPaymentRepository(assetRepo)
	repository.records = map[string]PaymentRecord{
		"1": createPaymentRecord(payment.Invest, 2000, "1"),
		"2": createPaymentRecord(payment.Return, 2019, "1"),
		"3": createPaymentRecord(payment.Invest, 2019, "2"),
		"4": createPaymentRecord(payment.Invest, 2019, "1"),
	}

	res, err := repository.FindByAssetNames(
		[]string{gold},
		[]payment.Period{payment.PeriodMock{
			TimeFrom:  payment.CreateYearDate(2018),
			TimeUntil: payment.CreateYearDate(2020),
		}},
		[]payment.Type{payment.Invest},
	)
	if err != nil {
		t.Errorf("Unexpeted err: %+v", err)
	}
	if len(res) != 1 {
		t.Errorf("Unepxted result, 1 item expected but got %d", len(res))
	}

	res, err = repository.FindByAssetNames(
		[]string{},
		[]payment.Period{},
		[]payment.Type{},
	)
	if err != nil {
		t.Errorf("Unexpeted err: %+v", err)
	}
	if len(res) != 4 {
		t.Errorf("All 4 items exected to be returned when no filters passed")
	}
}
