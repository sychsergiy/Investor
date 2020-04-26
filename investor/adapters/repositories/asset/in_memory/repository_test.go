package in_memory

import (
	"investor/entities/asset"
	"testing"
)

func TestInMemoryAssetRepository_Create(t *testing.T) {
	repository := NewInMemoryAssetRepository()
	a := asset.Asset{Id: "1", Category: asset.PreciousMetal, Name: "gold"}

	// save first payment, no errors expected
	err := repository.Create(a)
	if err != nil {
		t.Errorf("Unepxected error during payment creation: %s", err)
	}

	// try to save payment with the same id
	err = repository.Create(a)
	expectedErr := AssetAlreadyExistsError{"1"}
	if err != expectedErr {
		t.Error("Payment with id already exists error expected")
	}
}
