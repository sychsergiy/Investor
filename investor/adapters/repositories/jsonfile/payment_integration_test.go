package jsonfile

import (
	"errors"
	"investor/adapters/repositories/in_memory"
	"investor/entities/asset"
	"investor/entities/payment"
	"investor/helpers/file"
	"testing"
)

func checkErr(t *testing.T, err error, message string) bool {
	if err != nil {
		t.Errorf("Unexpected err %s %+v:", message, err)
		return true
	}
	return false
}

func TestPaymentRepository_Integration_ListAll(t *testing.T) {
	jsonFile := file.NewJsonFile(file.NewPlainFile(file.GetFilePath("test_list_all_payments.json")))
	storage := NewStorage(jsonFile)
	assetRepo := NewAssetRepository(storage)
	repo := NewPaymentRepository(storage, assetRepo)

	err := assetRepo.Create(asset.NewPlainAsset("assetId", asset.PreciousMetal, "name"))
	checkErr(t, err, "asset creation")

	err = repo.CreateBulk([]payment.Payment{
		payment.CreatePaymentWithAsset("1", "assetId", 2015),
		payment.CreatePaymentWithAsset("2", "assetId", 2016),
		payment.CreatePaymentWithAsset("3", "assetId", 2017),
	})
	checkErr(t, err, "payment bulk creation")

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
	err = repo2.Create(payment.CreatePaymentWithAsset("4", "assetId", 2018))
	checkErr(t, err, "payment creation")
	payments2, err = repo2.ListAll()
	checkErr(t, err, "payments list")
	if len(payments2) != 4 {
		t.Errorf("4 payments expected")
	}

	// test create work before first restoring
	repo3 := NewPaymentRepository(storage, assetRepo)
	err = repo3.Create(payment.CreatePaymentWithAsset("5", "assetId", 2019))
	checkErr(t, err, "payment creation")
	payments3, err := repo3.ListAll()
	checkErr(t, err, "payments list")

	if len(payments3) != 5 {
		t.Errorf("5 payments expected")
	}

	// test create returns error on none existent asset id
	p := in_memory.NewPaymentProxyMock(
		payment.CreatePayment("1", 2020),
		func() (a asset.Asset, err error) {
			return a, in_memory.AssetDoesntExistsError{AssetId: "not_exists"}
		},
	)
	err = repo3.Create(p)
	if err != nil {
		expectedErr := in_memory.AssetDoesntExistsError{AssetId: "not_exists"}
		if !errors.Is(err, expectedErr) {
			t.Errorf("AssetDoesntExistsError error expected, but got %s", err)
		}
	} else {
		t.Errorf("Asset with provided id doesnt exists error expected")
	}
}
