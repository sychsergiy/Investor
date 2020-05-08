package payment

import (
	"investor/entities/asset"
	"time"
)

func CreatePayment(id string, year int) Payment {
	testAsset := asset.NewPlainAsset("test", asset.CryptoCurrency, "test")
	creationTime := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	return NewPlainPayment(id, 0, 0, testAsset, creationTime, Invest)
}

func CreatePaymentWithAsset(id, assetId string, year int) Payment {
	testAsset := asset.NewPlainAsset(assetId, asset.CryptoCurrency, "test")
	creationTime := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	return NewPlainPayment(id, 0, 0, testAsset, creationTime, Invest)
}
