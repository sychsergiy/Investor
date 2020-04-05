package fiat

type Currency int

const (
	UAH Currency = iota
	USD
	EUR
)

func (d Currency) String() string {
	return [...]string{"UAH", "USD", "EUR"}[d]
}

type Rate float32
