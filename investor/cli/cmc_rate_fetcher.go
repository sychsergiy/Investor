package cli

import "investor/entities/asset"

// amount of USD per unit asset
type Rate float32

type RateFetcher interface {
	Fetch(asset.Asset) (Rate, error)
}
