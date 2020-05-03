package in_memory

import (
	"investor/entities/asset"
	"testing"
)

func TestInMemoryAssetRepository_Create(t *testing.T) {
	repository := NewAssetRepository()
	a := asset.Asset{Id: "1", Category: asset.PreciousMetal, Name: "gold"}

	// save first payment, no errors expected
	err := repository.Create(a)
	if err != nil {
		t.Errorf("Unepxected error during payment creation: %s", err)
	}

	// try to save payment with the same id
	err = repository.Create(a)
	expectedErr := RecordAlreadyExistsError{RecordId: "1"}
	if err != expectedErr {
		t.Error("Payment with id already exists error expected")
	}
}

func TestAssetRepository_CreateBulk(t *testing.T) {
	a1 := asset.Asset{Id: "1", Category: asset.CryptoCurrency, Name: "test"}
	a2 := asset.Asset{Id: "2", Category: asset.CryptoCurrency, Name: "test"}

	repository := NewAssetRepository()

	createdQuantity, err := repository.CreateBulk([]asset.Asset{a1, a2})
	if err != nil {
		t.Errorf("Unpected error")
		if createdQuantity != 2 {
			t.Errorf("2 payments expected to be created")
		}
	}

	repository = NewAssetRepository()
	expectedErr := RecordAlreadyExistsError{RecordId: "1"}
	createdQuantity, err = repository.CreateBulk([]asset.Asset{a1, a1})
	if err != expectedErr {
		t.Errorf("Payment alread exists error expected")
	}
	if createdQuantity != 1 {
		t.Errorf("One payment expected to be created before error")
	}
}
