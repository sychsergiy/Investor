package main

import (
	"investor/adapters"
	"investor/adapters/repositories/jsonfile"
	paymentRepo "investor/adapters/repositories/jsonfile/payment"
	"investor/cli/payment"
	"investor/cli/payment/rate_fetcher"
	"investor/helpers/file"
	"investor/interactors"
	"log"
	"os"
)

func setupDependencies(coinMarketCupApiKey string) payment.ConsolePaymentCreator {

	coinMarketCupClient := rate_fetcher.NewCoinMarketCupClient(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", coinMarketCupApiKey,
	)
	fetcher := rate_fetcher.CMCRateFetcher{
		Client: coinMarketCupClient,
	}

	jsonRepo := jsonfile.NewRepository(
		file.NewJsonFile(file.NewPlainFile("payments.json")),
		paymentRepo.Unmarshaler{},
	)
	repo := paymentRepo.NewRepository(jsonRepo)
	paymentCreateInteractor := interactors.CreatePayment{Repository: repo, IdGenerator: adapters.NewStubIdGenerator()}

	return payment.ConsolePaymentCreator{PaymentCreator: paymentCreateInteractor, RateFetcher: fetcher}
}
func main() {
	apiKey := os.Getenv("COIN_MARKET_CAP_API_KEY")
	if apiKey == "" {
		log.Fatal("COIN_MARKET_CAP_API_KEY env var not provided")
	}
	paymentCreator := setupDependencies(apiKey)
	err := paymentCreator.Create()
	if err != nil {
		log.Fatal(err)
	}
}
