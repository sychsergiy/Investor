package in_memory

import (
	"investor/entities/asset"
	paymentEntity "investor/entities/payment"
	"time"
)

func CreateAssetRecord(id string, name string) AssetRecord {
	return AssetRecord{id, asset.PreciousMetal, name}
}

func CreatePaymentRecord(id string, year int) PaymentRecord {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	return PaymentRecord{
		Id:             id,
		AssetAmount:    50,
		AbsoluteAmount: 100,
		AssetId:        "testAssetId",
		Type:           0,
		CreationDate:   date,
	}
}

type AssetFinderMock struct {
	findFunc func(assetId string) (*asset.Asset, error)
}

func (asm AssetFinderMock) FindById(assetId string) (*asset.Asset, error) {
	return asm.findFunc(assetId)
}

type PaymentProxyMock struct {
	paymentEntity.Payment
	assetFunc func() (*asset.Asset, error)
}

func (ppm PaymentProxyMock) Asset() (*asset.Asset, error) {
	return ppm.assetFunc()
}
