package asset

import "fmt"

type Category int

const (
	PreciousMetal Category = iota
	CryptoCurrency
	Stock
)

type AssetDoesntExistsError struct {
	AssetId string
}

func (e AssetDoesntExistsError) Error() string {
	return fmt.Sprintf("asset with id %s doesn't exist", e.AssetId)
}

func (c Category) String() string {
	return [...]string{"PreciousMetal", "CryptoCurrency", "Stock"}[c]
}

type Asset interface {
	Id() string
	Category() Category
	Name() string
}
