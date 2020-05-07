package jsonfile

import (
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
	jsonFile := file.NewJsonFile(file.NewPlainFile(file.GetFilePath("test_list_all.json")))
	repo := NewPaymentRepository(*NewStorage(jsonFile))

	err := repo.CreateBulk([]payment.Payment{
		payment.CreatePayment("1", 2015),
		payment.CreatePayment("2", 2016),
		payment.CreatePayment("3", 2017),
	})
	checkErr(t, err, "payment bulk creation")

	// test list in the same session works
	payments, err := repo.ListAll()
	checkErr(t, err, "payments list")
	if len(payments) != 3 {
		t.Errorf("3 payments expected")
	}

	// test restore from existent storage
	repo2 := NewPaymentRepository(*NewStorage(jsonFile))
	payments2, err := repo2.ListAll()
	checkErr(t, err, "payments list")
	if len(payments2) != 3 {
		t.Errorf("3 payments expected")
	}

	// test create works after restore (restored with first ListAll() call)
	err = repo2.Create(payment.CreatePayment("4", 2018))
	checkErr(t, err, "payment creation")
	payments2, err = repo2.ListAll()
	checkErr(t, err, "payments list")
	if len(payments2) != 4 {
		t.Errorf("4 payments expected")
	}

	// test create work before first restoring
	repo3 := NewPaymentRepository(*NewStorage(jsonFile))
	err = repo3.Create(payment.CreatePayment("5", 2019))
	checkErr(t, err, "payment creation")
	payments3, err := repo3.ListAll()
	checkErr(t, err, "payments list")

	if len(payments3) != 5 {
		t.Errorf("5 payments expected")
	}
}
