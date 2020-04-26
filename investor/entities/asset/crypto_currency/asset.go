package crypto_currency

import (
	"investor/entities/asset"
)

type CryptoCurrency string

const (
	BTC  CryptoCurrency = "BTC"
	ETH                 = "ETH"
	XRP                 = "XRP"
	DASH                = "DASH"
)

func NewAsset(id string, name CryptoCurrency) asset.Asset {
	return asset.Asset{Id: id, Category: asset.CryptoCurrency, Name: string(name)}
}
