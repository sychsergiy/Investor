package crypto_currency

import "investor/asset"

type CryptoCurrency string

const (
	BTC  CryptoCurrency = "BTC"
	ETH                 = "ETH"
	XRP                 = "XRP"
	DASH                = "DASH"
)

func NewAsset(name CryptoCurrency) asset.Asset {
	return asset.Asset{Category: asset.CryptoCurrency, Name: string(name)}
}
