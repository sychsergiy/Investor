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
	type_          Type
}

func (p PlainPayment) Id() string {
	return p.id
}

func (p PlainPayment) AssetAmount() float32 {
	return p.assetAmount
}

func (p PlainPayment) AbsoluteAmount() float32 {
	return p.absoluteAmount
}

func (p PlainPayment) Asset() asset.Asset {
	return p.asset
}

func (p PlainPayment) CreationDate() time.Time {
	return p.creationDate
}

func (p PlainPayment) Type() Type {
	return p.type_
}

func NewPlainPayment(
	id string, assetAmount float32, absoluteAmount float32,
	asset asset.Asset, creationDate time.Time, type_ Type,
) PlainPayment {
	return PlainPayment{id, assetAmount, absoluteAmount, asset, creationDate, type_}
}
