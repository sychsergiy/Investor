package crypto

type Currency int

const (
	BTC  Currency = iota // Bitcoin
	ETH                  // Ethereum
	Dash                 // Dash
	XRP                  // XRP
)

func (d Currency) String() string {
	return [...]string{"BTC", "ETH", "Dash", "XRP"}[d]
}
