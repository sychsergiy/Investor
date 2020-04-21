package payment

import (
	"investor/asset"
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

type Payment interface {
	AbsoluteAmount() float32
	AssetAmount() float32
	Asset() asset.Asset
	Type() Type
	CreationDate() time.Time
}
