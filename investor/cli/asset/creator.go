package asset

import (
	"bufio"
	"fmt"
	"investor/entities/asset"
	"investor/interactors"
	"os"
	"strconv"
)

type ConsoleAssetCreator struct {
	creator interactors.CreateAsset
}

func NewConsoleAssetCreator(creator interactors.CreateAsset) ConsoleAssetCreator {
	return ConsoleAssetCreator{creator}
}

func (c ConsoleAssetCreator) Execute() {
	category := chooseCategory()
	name := readName()

	request := interactors.CreateAssetRequest{Name: name, Category: category}
	response := c.creator.Create(request)

	if response.Created {
		fmt.Println("OK. Asset created")
	} else {
		fmt.Printf("Failed to created asset due to err %s\n", response.Err)
	}
}

func readFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func chooseCategory() asset.Category {
	fmt.Printf(
		"Choose category: %d - Precious metal, %d - Crypto Currency, %d - Stock\n",
		asset.PreciousMetal, asset.CryptoCurrency, asset.Stock,
	)
	input := readFromConsole()

	number, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("Unexpected input, failed due to error: %s\n", err))
	}

	return asset.Category(number)
}

func readName() string {
	fmt.Println("Enter an asset name:")
	input := readFromConsole()
	return input
}
