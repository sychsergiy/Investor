package payment

import (
	"investor/entities/asset"
	"time"
)

func CreatePayment(id string, year int) Payment {
	testAsset := asset.Asset{Id: "test", Category: asset.CryptoCurrency, Name: "test"}
	creationTime := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	return NewReturn(id, 0, 0, testAsset, creationTime)
}

func CreatePaymentWithAsset(id, assetId string, year int) Payment {
	testAsset := asset.Asset{Id: assetId, Category: asset.CryptoCurrency, Name: "test"}
	creationTime := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	return NewReturn(id, 0, 0, testAsset, creationTime)
}
