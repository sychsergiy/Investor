package asset

import "fmt"

type Category int

const (
	PreciousMetal Category = iota
	CryptoCurrency
	Stock
)

type NotFoundError struct {
	AssetID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("asset with id %s doesn't exist", e.AssetID)
}

func (c Category) String() string {
	return [...]string{"PreciousMetal", "CryptoCurrency", "Stock"}[c]
}

type Asset interface {
	ID() string
	Category() Category
	Name() string
}
