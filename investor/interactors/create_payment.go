package interactors

import (
	"bufio"
	"fmt"
	"investor/entities"
	"investor/entities/asset"
	"investor/entities/asset/crypto_currency"
	"investor/ports"
	"os"
	"strconv"
	"time"
)

type PaymentCreator struct {
	Storage     ports.PaymentSaver
	RateFetcher ports.RateFetcher
	IdGenerator ports.IdGenerator
}

func (pc PaymentCreator) Create() {
	paymentType := choosePaymentType()
	asset_ := chooseAsset()
	date := readCreationDate()

	rate, err := pc.RateFetcher.Fetch(asset_)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Suggested currency rate %f\n", rate) // todo: suggest for creation date, now today

	fmt.Println("Enter invested amount in USD: ")
	absoluteAmount := readAmount()

	fmt.Printf("Suggested asset amount %f\n", float32(rate)*absoluteAmount)
	fmt.Println("Enter invested amount in asset: ")
	assetAmount := readAmount()

	var payment entities.Payment
	id := pc.IdGenerator.Generate()
	if paymentType == entities.Return {
		payment = entities.NewReturnPayment(id, assetAmount, absoluteAmount, asset_, date)
	} else if paymentType == entities.Invest {
		payment = entities.NewInvestmentPayment(id, assetAmount, absoluteAmount, asset_, date)
	} else {
		panic(fmt.Sprintf("unexpected ports type: %d", paymentType))
	}

	saveRecord := readCompleteOrAbort(payment)
	if saveRecord {
		pc.Storage.Save(payment)
	}
}

func readFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func readCompleteOrAbort(record entities.Payment) bool {
	fmt.Println("Verify created ports: ")
	fmt.Printf("%+v\n", record)
	fmt.Println("Enter:  1 - to save record, 2 - to abort")
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
	layout := "01-02-2006 15:04:05"
	fmt.Println("Enter creation date in the following format: dd-mm-yyyy hh:mm:ss")
	input := readFromConsole()
	creationDate, err := time.Parse(layout, input)
	if err != nil {
		panic(err)
	}
	return creationDate
}

func ParseTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

func choosePaymentType() entities.Type {
	fmt.Println("Choose ports type:\n 1 - Invest \n 2 - Return")
	switch input := readFromConsole(); input {
	case "1":
		return entities.Invest
	case "2":
		return entities.Return
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
	return crypto_currency.NewAsset(currency)
}
