package payment

import (
	"bufio"
	"fmt"
	"investor/cli/payment/rate_fetcher"
	"investor/entities/asset"
	"investor/entities/asset/crypto_currency"
	"investor/entities/payment"
	"investor/interactors"
	"os"
	"strconv"
	"time"
)

type ConsolePaymentCreator struct {
	PaymentCreator interactors.PaymentCreator
	RateFetcher    rate_fetcher.RateFetcher
}

func (cpc ConsolePaymentCreator) Create() error {
	paymentType := choosePaymentType()
	asset_ := chooseAsset()
	date := readCreationDate()

	rate, err := cpc.RateFetcher.Fetch(asset_)
	if err != nil {
		return err
	}
	fmt.Printf("Suggested currency rate %f\n", rate) // todo: suggest for creation date, now today

	fmt.Println("Enter invested amount in USD: ")
	absoluteAmount := readAmount()

	fmt.Printf("Suggested asset amount %f\n", float32(rate)*absoluteAmount)
	fmt.Println("Enter invested amount in asset: ")
	assetAmount := readAmount()

	model := interactors.CreatePaymentModel{
		AssetAmount: assetAmount, AbsoluteAmount: absoluteAmount,
		Asset: asset_, Type: paymentType, CreationDate: date,
	}
	saveRecord := readCompleteOrAbort(model)
	if saveRecord {
		err = cpc.PaymentCreator.Create(model)
	}
	return err
}

func readFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func readCompleteOrAbort(model interactors.CreatePaymentModel) bool {
	fmt.Println("Verify created payment: ")
	fmt.Printf("%+v\n", model)
	fmt.Println("Enter:  1 - to save model, 2 - to abort")
	input := readFromConsole()
	if input == "1" {
		return true
	} else if input == "2" {
		return false
	} else {
		panic(fmt.Sprintf("Unexpected input: %s", input))
	}
}

func readCreationDate() time.Time {
	fmt.Println("Enter creation date in the following format: yyyy-dd-mm hh:mm:ss")
	input := readFromConsole()
	creationDate, err := ParseTime(input)
	if err != nil {
		panic(err)
	}
	return creationDate
}

func ParseTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

func choosePaymentType() payment.Type {
	fmt.Println("Choose adapters type:\n 1 - Invest \n 2 - Return")
	switch input := readFromConsole(); input {
	case "1":
		return payment.Invest
	case "2":
		return payment.Return
	default:
		panic("Unexpected input")
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

func chooseAsset() asset.Asset {
	fmt.Println("Choose crypto currency:\n 1 - BTC(bitcoin) \n 2 - ETH(ethereum) \n 3 - Dash \n 4 - XRP")
	var currency crypto_currency.CryptoCurrency
	switch input := readFromConsole(); input {
	case "1":
		currency = crypto_currency.BTC
	case "2":
		currency = crypto_currency.ETH
	case "3":
		currency = crypto_currency.DASH
	case "4":
		currency = crypto_currency.XRP
	default:
		panic("Unexpected input")
	}
	return crypto_currency.NewAsset("test", currency)
}
