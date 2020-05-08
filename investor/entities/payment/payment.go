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

type Payment interface {
	Id() string
	AssetAmount() float32
	AbsoluteAmount() float32
	Asset() (*asset.Asset, error)
	Type() Type
	CreationDate() time.Time
}
