package jsonfile

import (
	"errors"
	"investor/adapters/repositories/memory"
	"investor/entities/asset"
	"investor/entities/payment"
	"reflect"
	"testing"
)

func checkErr(t *testing.T, err error, message string) bool {
	if err != nil {
		t.Errorf("Unexpected err %s %+v:", message, err)
		return true
	}
	return false
}

func FillPaymentsRepo(t *testing.T, paymentRepo *PaymentRepository, assetRepo *AssetRepository) {
	assetID := "assetID"
	err := assetRepo.Create(asset.NewPlainAsset(assetID, asset.PreciousMetal, "name"))
	checkErr(t, err, "asset creation")

	err = paymentRepo.CreateBulk([]payment.Payment{
		payment.CreatePaymentWithAsset("1", assetID, 2015),
		payment.CreatePaymentWithAsset("2", assetID, 2016),
		payment.CreatePaymentWithAsset("3", assetID, 2017),
	})
	checkErr(t, err, "payment bulk creation")
}

func TestPaymentRepository_Integration_ListAll(t *testing.T) {
	storage := createStorage("test_list_all_payments.json")
	assetRepo := NewAssetRepository(storage)
	repo := NewPaymentRepository(storage, assetRepo)

	FillPaymentsRepo(t, repo, assetRepo)

	// test list in the same session works
	payments, err := repo.ListAll()
	checkErr(t, err, "payments list")
	if len(payments) != 3 {
		t.Errorf("3 payments expected")
	}

	// test restore from existent storage
	repo2 := NewPaymentRepository(storage, assetRepo)
	payments2, err := repo2.ListAll()
	checkErr(t, err, "payments list")
	if len(payments2) != 3 {
		t.Errorf("3 payments expected")
	}

	// test create works after restore (restored with first ListAll() call)
	err = repo2.Create(payment.CreatePaymentWithAsset("4", "assetID", 2018))
	checkErr(t, err, "payment creation")
	payments2, err = repo2.ListAll()
	checkErr(t, err, "payments list")
	if len(payments2) != 4 {
		t.Errorf("4 payments expected")
	}

	// test create work before first restoring
	repo3 := NewPaymentRepository(storage, assetRepo)
	err = repo3.Create(payment.CreatePaymentWithAsset("5", "assetID", 2019))
	checkErr(t, err, "payment creation")
	payments3, err := repo3.ListAll()
	checkErr(t, err, "payments list")

	if len(payments3) != 5 {
		t.Errorf("5 payments expected")
	}

	// test create returns error on none existent asset id
	p := payment.NewProxyMock(
		payment.CreatePayment("1", 2020),
		func() (a asset.Asset, err error) {
			return a, asset.NotFoundError{AssetID: "not_exists"}
		},
	)
	err = repo3.Create(p)
	if err != nil {
		expectedErr := asset.NotFoundError{AssetID: "not_exists"}
		if !errors.Is(err, expectedErr) {
			t.Errorf("NotFoundError error expected, but got %s", err)
		}
	} else {
		t.Errorf("Asset with provided id doesnt exists error expected")
	}
}

func TestPaymentRepository_Integration_FindByIDs(t *testing.T) {
	storage := createStorage("test_payments_by_ids.json")
	assetRepo := NewAssetRepository(storage)
	repo := NewPaymentRepository(storage, assetRepo)

	FillPaymentsRepo(t, repo, assetRepo)

	payments, err := repo.FindByIDs([]string{"1", "3"})
	if err != nil {
		t.Errorf("Unexpected err %+v", err)
	} else {
		expectedIDs := []string{"3", "1"}
		ids := payment.PaymentsToIDs(payments)
		if !reflect.DeepEqual(ids, expectedIDs) {
			t.Errorf("Unexpected payment ids")
		}
	}

	payments, err = repo.FindByIDs([]string{"not_existent"})
	if !errors.Is(err, memory.PaymentDoesntExistsError{PaymentID: "not_existent"}) {
		t.Errorf("Unexpected err %+v", err)
	}
	if payments != nil {
		t.Errorf("Payments nil value expected")
	}

}
