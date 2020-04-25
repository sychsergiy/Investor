package entities

import (
	"investor/entities/asset"
	"time"
)

type PaymentType int

const (
	Invest PaymentType = iota
	Return
)

func (t PaymentType) String() string {
	return [...]string{"Invest", "Return"}[t]
}

type Payment struct {
	Id             string
	AssetAmount    float32
	AbsoluteAmount float32
	Asset          asset.Asset
	Type           PaymentType
	CreationDate   time.Time
}

func NewInvestmentPayment(
	id string, assetAmount float32, absoluteAmount float32,
	asset asset.Asset, creationDate time.Time,
) Payment {
	return Payment{
		Id:             id,
		AssetAmount:    assetAmount,
		AbsoluteAmount: absoluteAmount,
		Asset:          asset,
		Type:           Invest,
		CreationDate:   creationDate,
	}
}

func NewReturnPayment(
	id string, assetAmount float32, absoluteAmount float32,
	asset asset.Asset, creationDate time.Time,
) Payment {
	return Payment{
		Id:             id,
		AssetAmount:    assetAmount,
		AbsoluteAmount: absoluteAmount,
		Asset:          asset,
		Type:           Return,
		CreationDate:   creationDate,
	}
}
