package asset

import (
	"Investor/asset/amount"
	"Investor/asset/amount/crypto"
	"Investor/asset/amount/fiat"
	"Investor/asset/payment"
	"Investor/asset/payment/storage"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func MakeInvestment(storage storage.PaymentsSaver, invested amount.Amount, cryptoInvestedAmount float32, date time.Time) {
	p := payment.NewInvestPayment(invested.AbsoluteValue(), cryptoInvestedAmount, date)
	storage.SavePayment(p)
}

func MakeReturn(storage storage.PaymentsSaver, returned amount.Amount, cryptoReturnedAmount float32, date time.Time) {
	p := payment.NewReturnPayment(returned.AbsoluteValue(), cryptoReturnedAmount, date)
	storage.SavePayment(p)
}

func CreateInvestment() {
	// choose fiat Currency
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose currency 1 - UAH, 2 - USD: ")
	scanner.Scan()
	var currency fiat.Currency
	switch input := scanner.Text(); input {
	case "1":
		currency = fiat.UAH
	case "2":
		currency = fiat.USD
	default:
		log.Fatal("Not expected input")
	}

	if scanner.Err() != nil {
		log.Fatal("scanner error")
	}

	// suggest USD amount
	// maybe manual entering
	var currencyRate float32
	var currencyConvertRequired = false
	if currency == fiat.UAH {
		currencyConvertRequired = true
		suggest := FetchUAHCurrencyRate()
		fmt.Printf("UAH to USD currency rate: %f \n", suggest)
		fmt.Printf("Use suggested(1) or manual entering(2): ")

		scanner.Scan()

		switch input := scanner.Text(); input {
		case "1":
			currencyRate = suggest
		case "2":
			scanner.Scan()
			rateInput := scanner.Text()
			number, err := strconv.ParseFloat(rateInput, 32)
			if err != nil {
				log.Fatal(err)
			}
			currencyRate = float32(number)
		default:
			log.Fatal("Not expected input")
		}

		if scanner.Err() != nil {
			log.Fatal("scanner error")
		}
	}
	fmt.Println(currencyRate)

	// input invested amount
	fmt.Println("Please enter amount: ")
	scanner.Scan()
	amountInput := scanner.Text()

	amount_, err := strconv.ParseFloat(amountInput, 32)
	if err != nil {
		log.Fatal(err)
	}
	investedAmount := float32(amount_)

	if currencyConvertRequired {
		investedAmount /= currencyRate
	}

	// choose crypto Currency
	cryptoCurrency := chooseCryptoCurrency()

	// suggest crypto crypto Amount
	cryptoCurrencyRate := fetchCryptoCurrencyRate(cryptoCurrency, os.Getenv("COIN_MARKET_CAP_API_KEY"))
	suggestCryptoAmount := investedAmount / cryptoCurrencyRate

	// maybe manual entering
	fmt.Printf(
		"Approximately invested sum: %f %s with rate: %f USD\n",
		suggestCryptoAmount, cryptoCurrency, cryptoCurrencyRate)

	fmt.Printf("Use suggested(1) or manual entering(2): ")
	scanner.Scan()
	var investedCryptoSum float32
	switch input := scanner.Text(); input {
	case "1":
		investedCryptoSum = suggestCryptoAmount
	case "2":
		investedCryptoSum = inputInvestedCryptoSum(scanner)
	}

	fmt.Printf("Invested crypto sum %f\n", investedCryptoSum)

	// input date
}

func inputInvestedCryptoSum(scanner *bufio.Scanner) float32 {
	fmt.Println("Enter invested sum in crypto currency: ")
	scanner.Scan()
	input := scanner.Text()
	investedSum_, err := strconv.ParseFloat(input, 32)
	if err != nil {
		log.Fatal(err)
	}
	return float32(investedSum_)
}

func chooseCryptoCurrency() crypto.Currency {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose crypto currency:\n 1 - BTC(bitcoin) \n 2 - ETH(ethereum) \n 3 - Dash \n 4 - XRP")
	scanner.Scan()
	switch input := scanner.Text(); input {
	case "1":
		return crypto.BTC
	case "2":
		return crypto.ETH
	case "3":
		return crypto.Dash
	case "4":
		return crypto.XRP
	default:
		panic("Unexpected input")
	}
}