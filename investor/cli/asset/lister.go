package asset

import (
	"fmt"
	"investor/entities/asset"
	"investor/interactors"
)

type ConsoleAssetsLister struct {
	lister interactors.ListAssets
}

func (l ConsoleAssetsLister) Execute() {
	l.List()
}

func (l ConsoleAssetsLister) List() {
	assets, err := l.lister.ListAll()
	if err != nil {
		panic(fmt.Errorf("failed to list assets: %+v", err))
	}

	fmt.Printf("Total assets count: %d\n", len(assets))
	for i, p := range assets {
		str := AssetToString(p)
		fmt.Printf("%d -------------------------\n", i+1)
		println(str)
	}
}

func AssetToString(a asset.Asset) string {
	return fmt.Sprintf(
		"ID: %s\nName: %s\nCategory: %s\n",
		a.Id(), a.Name(), a.Category(),
	)

}

func NewConsoleAssetLister(lister interactors.ListAssets) ConsoleAssetsLister {
	return ConsoleAssetsLister{lister}
}
