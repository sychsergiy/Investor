package ports

import (
	"investor/entities/asset"
	"testing"
	"time"
)

func TestInMemoryStorage_Create(t *testing.T) {
	storage := NewInMemoryStorage()
	testAsset := asset.Asset{Category: asset.CryptoCurrency, Name: "test"}
	creationTime := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)
	id := storage.Create(NewReturn(0, 0, testAsset, creationTime))
	if id != Identifier(0) {
		t.Errorf("Identfier of first created payemnt should be 0", )
	}
	id = storage.Create(NewInvestment(0, 0, testAsset, creationTime))
	if id != Identifier(1) {
		t.Errorf("Identfier of second created ports should be 1", )
	}
}
