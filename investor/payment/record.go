package payment

import (
	"investor/asset"
	"time"
)

type Record struct {
	assetAmount    float32
	absoluteAmount float32
	asset          asset.Asset
	type_          Type
	creationDate   time.Time
}

func (r Record) AbsoluteAmount() float32 {
	return r.absoluteAmount
}

func (r Record) AssetAmount() float32 {
	return r.assetAmount
}

func (r Record) Type() Type {
	return r.type_
}

func (r Record) CreationDate() time.Time {
	return r.creationDate
}

func (r Record) Asset() asset.Asset {
	return r.asset
}

func NewInvestment(assetAmount float32, absoluteAmount float32, asset asset.Asset, creationDate time.Time) Record {
	return Record{
		assetAmount:    assetAmount,
		absoluteAmount: absoluteAmount,
		asset:          asset,
		type_:          Invest,
		creationDate:   creationDate,
	}
}

func NewReturn(assetAmount float32, absoluteAmount float32, asset asset.Asset, creationDate time.Time) Record {
	return Record{
		assetAmount:    assetAmount,
		absoluteAmount: absoluteAmount,
		asset:          asset,
		type_:          Return,
		creationDate:   creationDate,
	}
}
