package interactors

import (
	"investor/entities/asset"
	"investor/interactors/ports"
)

type ListAssets struct {
	lister ports.AssetsLister
}

func (lp ListAssets) ListAll() ([]asset.Asset, error) {
	return lp.lister.ListAll()
}

func NewListAssets(lister ports.AssetsLister) ListAssets {
	return ListAssets{lister}
}
