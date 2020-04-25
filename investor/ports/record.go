package ports

import (
	"investor/entities"
	asset2 "investor/entities/asset"
	"time"
)

type Record struct {
	assetAmount    float32
	absoluteAmount float32
	asset          asset2.Asset
	type_          entities.Type
	creationDate   time.Time
}

func (r Record) AbsoluteAmount() float32 {
	return r.absoluteAmount
}

func (r Record) AssetAmount() float32 {
	return r.assetAmount
}

func (r Record) Type() entities.Type {
	return r.type_
}

func (r Record) CreationDate() time.Time {
	return r.creationDate
}

func (r Record) Asset() asset2.Asset {
	return r.asset
}

func NewInvestment(assetAmount float32, absoluteAmount float32, asset asset2.Asset, creationDate time.Time) Record {
	return Record{
		assetAmount:    assetAmount,
		absoluteAmount: absoluteAmount,
		asset:          asset,
		type_:          entities.Invest,
		creationDate:   creationDate,
	}
}

func NewReturn(assetAmount float32, absoluteAmount float32, asset asset2.Asset, creationDate time.Time) Record {
	return Record{
		assetAmount:    assetAmount,
		absoluteAmount: absoluteAmount,
		asset:          asset,
		type_:          entities.Return,
		creationDate:   creationDate,
	}
}
