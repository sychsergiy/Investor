package asset

import (
	"fmt"
	"investor/entities/asset"
	"investor/interactors"
)

type ListAssetsCommand struct {
	lister interactors.ListAssets
}

func (l ListAssetsCommand) Execute() {
	l.List()
}

func (l ListAssetsCommand) List() {
	assets, err := l.lister.ListAll()
	if err != nil {
		panic(fmt.Errorf("failed to list assets: %+v", err))
	}

	fmt.Printf("Total assets count: %d\n", len(assets))
	for i, p := range assets {
		str := ConvertAssetToString(p)
		fmt.Printf("%d -------------------------\n", i+1)
		println(str)
	}
}

func ConvertAssetToString(a asset.Asset) string {
	return fmt.Sprintf(
		"ID: %s\nName: %s\nCategory: %s\n",
		a.ID(), a.Name(), a.Category(),
	)

}

func NewListAssetsCommand(lister interactors.ListAssets) ListAssetsCommand {
	return ListAssetsCommand{lister}
}
