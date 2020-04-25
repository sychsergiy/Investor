package payment

import (
	"investor/entities/asset"
	"time"
)

type Type int

const (
	Invest Type = iota
	Return
)

func (t Type) String() string {
	return [...]string{"Invest", "Return"}[t]
}

type Payment struct {
	Id             string
	AssetAmount    float32
	AbsoluteAmount float32
	Asset          asset.Asset
	Type           Type
	CreationDate   time.Time
}

func NewInvestment(
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

func NewReturn(
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
