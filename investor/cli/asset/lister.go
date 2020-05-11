package asset

import (
	"fmt"
	"investor/entities/asset"
	"investor/interactors"
)

type ListCommand struct {
	lister interactors.ListAssets
}

func (l ListCommand) Execute() {
	l.List()
}

func (l ListCommand) List() {
	assets, err := l.lister.ListAll()
	if err != nil {
		panic(fmt.Errorf("failed to list assets: %+v", err))
	}

	fmt.Printf("Total assets count: %d\n", len(assets))
	for i, p := range assets {
		str := ToString(p)
		fmt.Printf("%d -------------------------\n", i+1)
		println(str)
	}
}

func ToString(a asset.Asset) string {
	return fmt.Sprintf(
		"ID: %s\nName: %s\nCategory: %s\n",
		a.ID(), a.Name(), a.Category(),
	)

}

func NewListCommand(lister interactors.ListAssets) ListCommand {
	return ListCommand{lister}
}
