package payment

import (
	"bufio"
	"fmt"
	assetCLI "investor/cli/asset"
	"investor/cli/payment/rate"
	"investor/entities/asset"
	"investor/entities/payment"
	"investor/interactors"
	"log"
	"os"
	"strconv"
	"time"
)

type CreatePaymentCommand struct {
	paymentCreator interactors.CreatePayment
	assetsLister   interactors.ListAssets
	rateFetcher    rate.Fetcher
}

func NewCreatePaymentCommand(
	paymentCreator interactors.CreatePayment,
	assetsLister interactors.ListAssets,
	rateFetcher rate.Fetcher,
) CreatePaymentCommand {
	return CreatePaymentCommand{paymentCreator, assetsLister, rateFetcher}
}

func (cpc CreatePaymentCommand) Execute() {
	err := cpc.Create()
	if err != nil {
		log.Fatal(err)
	}
}
func (cpc CreatePaymentCommand) Create() error {
	paymentType := choosePaymentType()
	//selectedAsset := chooseAsset()
	selectedAsset := cpc.selectAsset()
	date := readCreationDate()

	//rate, err := cpc.rateFetcher.Fetch(selectedAsset)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("Suggested currency rate %f\n", rate) // todo: suggest for creation date, now today

	fmt.Println("Enter invested amount in USD: ")
	absoluteAmount := readAmount()

	//fmt.Printf("Suggested asset amount %f\n", float32(rate)*absoluteAmount)
	fmt.Println("Enter invested amount in asset: ")
	assetAmount := readAmount()

	model := interactors.CreatePaymentModel{
		AssetAmount: assetAmount, AbsoluteAmount: absoluteAmount,
		Asset: selectedAsset, Type: paymentType, CreationDate: date,
	}
	saveRecord := readCompleteOrAbort(model)
	if saveRecord {
		return cpc.paymentCreator.Create(model)
	}
	fmt.Println("Aborted.")
	return nil
}

func readFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func readCompleteOrAbort(model interactors.CreatePaymentModel) bool {
	fmt.Println("Verify created payment Enter:  1 - to save model, 2 - to abort: ")
	fmt.Printf("----------------\n%s\n----------------\n", paymentModelToString(model))
	input := readFromConsole()
	switch input {
	case "1":
		return true
	case "2":
		return false
	default:
		log.Fatalf(fmt.Sprintf("Unexpected input: %s", input))
	}
	return false
}

func readCreationDate() time.Time {
	fmt.Println("Enter creation date in the following format: yyyy-mm-dd hh:mm")
	input := readFromConsole()
	creationDate, err := ParseTime(input)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return creationDate
}

func paymentModelToString(model interactors.CreatePaymentModel) string {
	str := fmt.Sprintf(
		"Asset Name: %s\nAsset category: %s\nType: %s\nUSD amount: %.2f\nAsset amount: %.5f\nCreation date: %s",
		model.Asset.Name(), model.Asset.Category(), model.Type,
		model.AbsoluteAmount, model.AssetAmount, model.CreationDate,
	)
	return str
}

func ParseTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04", str)
}

func choosePaymentType() payment.Type {
	fmt.Println("Choose payment type:\n 1 - Invest \n 2 - Return")
	switch input := readFromConsole(); input {
	case "1":
		return payment.Invest
	case "2":
		return payment.Return
	default:
		log.Fatal("Unexpected input")
		return 0
	}
}

func readAmount() float32 {
	input := readFromConsole()
	value, err := strconv.ParseFloat(input, 32)
	if err != nil {
		// do something sensible
	}
	float := float32(value)
	return float
}

func (cpc CreatePaymentCommand) selectAsset() asset.Asset {

	assets, err := cpc.assetsLister.ListAll()

	assetsLen := len(assets)

	if err != nil {
		log.Fatalf("Err :%+v", err)
	}

	fmt.Println("-------------Assets------------------")

	for i, p := range assets {
		if i != 0 && i != assetsLen {
			fmt.Println("-------------------------------", )

		}
		str := assetCLI.ConvertAssetToString(p)
		fmt.Printf("Number: %d\n", i+1)
		fmt.Println(str)
	}

	fmt.Println("---------------End-------------------")

	fmt.Printf("Enter a number from 1 to %d to select asset: ", assetsLen)

	input := readFromConsole()
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if number < 1 || number > assetsLen {
		log.Fatal("wrong input, out of assets range")
	}

	a := assets[number-1]

	str := assetCLI.ConvertAssetToString(a)
	fmt.Printf("Selected asset:\n%s\n", str)

	return a
}
