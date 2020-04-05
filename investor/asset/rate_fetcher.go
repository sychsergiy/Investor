package asset

// amount of USD per unit asset
type Rate float32

type RateFetcher interface {
	Fetch(Asset) (Rate, error)
}
