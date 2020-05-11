package rate

import "investor/entities/asset"

// amount of USD per unit asset
type Rate float32

type Fetcher interface {
	Fetch(asset.Asset) (Rate, error)
}
