package currency

type Currency int

const (
	UAH Currency = iota
	USD
)

func (d Currency) String() string {
	return [...]string{"UAH", "USD"}[d]
}
