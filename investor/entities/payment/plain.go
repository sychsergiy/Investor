package payment

import (
	"investor/entities/asset"
	"time"
)

type PlainPayment struct {
	id             string
	assetAmount    float32
	absoluteAmount float32
	asset          asset.Asset
	creationDate   time.Time
	paymentType    Type
}

func (p PlainPayment) ID() string {
	return p.id
}

func (p PlainPayment) AssetAmount() float32 {
	return p.assetAmount
}

func (p PlainPayment) AbsoluteAmount() float32 {
	return p.absoluteAmount
}

func (p PlainPayment) Asset() (asset.Asset, error) {
	return p.asset, nil
}

func (p PlainPayment) CreationDate() time.Time {
	return p.creationDate
}

func (p PlainPayment) Type() Type {
	return p.paymentType
}

func NewPlainPayment(
	id string, assetAmount float32, absoluteAmount float32,
	asset asset.Asset, creationDate time.Time, paymentType Type,
) PlainPayment {
	return PlainPayment{
		id:             id,
		assetAmount:    assetAmount,
		absoluteAmount: absoluteAmount,
		asset:          asset,
		creationDate:   creationDate,
		paymentType:    paymentType,
	}
}
