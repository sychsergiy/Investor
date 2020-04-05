package asset

type Category int

const (
	PreciousMetal Category = iota
	CryptoCurrency
	Stock
)

type Asset struct {
	Category string
	Name     string
}

type CryptoCurrencies string

const (
	BTC  CryptoCurrencies = "BTC"
	ETH                   = "ETH"
	XRP                   = "XRP"
	DASH                  = "DASH"
)

func NewCryptoCurrency(name CryptoCurrencies) Asset {
	return Asset{string(CryptoCurrency), string(name)}
}
