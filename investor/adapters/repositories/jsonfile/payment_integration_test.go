package jsonfile

import (
	"investor/entities/payment"
	"investor/helpers/file"
	"testing"
)

func TestPaymentRepository_Integration_ListAll(t *testing.T) {
	jsonFile := file.NewJsonFile(file.NewPlainFile(file.GetFilePath("test_list_all.json")))
	repo := NewPaymentRepository(*NewStorage(jsonFile))

	_, err := repo.CreateBulk([]payment.Payment{
		payment.CreatePayment("1", 2015),
		payment.CreatePayment("2", 2016),
		payment.CreatePayment("3", 2017),
	})
	if err != nil {
		t.Errorf("Unexpected payments creation err %+v", err)
	}

	// test list in the same session works
	payments := repo.ListAll()
	if len(payments) != 3 {
		t.Errorf("3 payments expected")
	}

	// test restore from existent storage
	repo2 := NewPaymentRepository(*NewStorage(jsonFile))
	payments2 := repo2.ListAll()
	if len(payments2) != 3 {
		t.Errorf("3 payments expected")
	}

	// test create works after restore (restored with first ListAll() call)
	err = repo2.Create(payment.CreatePayment("4", 2018))
	if err != nil {
		t.Errorf("Unexpected payments creation err %+v", err)
	}
	payments2 = repo2.ListAll()
	if len(payments2) != 4 {
		t.Errorf("4 payments expected")
	}

	// test create work before first restoring
	repo3 := NewPaymentRepository(*NewStorage(jsonFile))
	err = repo3.Create(payment.CreatePayment("5", 2019))
	if err != nil {
		t.Errorf("Unexpected payment creation err %+v", err)
	}
	payments3 := repo3.ListAll()
	if len(payments3) != 5 {
		t.Errorf("5 payments expected")
	}
}
