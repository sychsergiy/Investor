package asset

import (
	"bufio"
	"fmt"
	"investor/entities/asset"
	"investor/interactors"
	"log"
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

	complete := readCompleteOrAbort(request)
	if complete {
		response := c.creator.Create(request)
		if response.Created {
			fmt.Println("OK. Asset created")
		} else {
			fmt.Printf("Failed to created asset due to err %s\n", response.Err)
		}
	} else {
		fmt.Println("Aborted.")
	}

}

func readCompleteOrAbort(model interactors.CreateAssetRequest) bool {
	fmt.Printf(
		"Verify asset. Enter:  1 - to save, 2 - to abort: \n------------\nName: %s\nCategory: %s\n------------\n",
		model.Name, model.Category.String(),
	)
	input := readFromConsole()
	if input == "1" {
		return true
	} else if input == "2" {
		return false
	} else {
		log.Fatal(fmt.Sprintf("Unexpected input: %s", input))
		return false
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
		log.Fatalf("Unexpected input, failed due to error: %s\n", err)
	}

	category := asset.Category(number)
	fmt.Printf("%s selected\n", category.String())
	return category
}

func readName() string {
	fmt.Println("Enter an asset name:")
	input := readFromConsole()
	return input
}
