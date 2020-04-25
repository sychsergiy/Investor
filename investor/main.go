package main

import (
	"investor/cli"
	"investor/interactors"
	"investor/adapters"
	"log"
	"os"
)

func setupDependencies(coinMarketCupApiKey string) cli.ConsolePaymentCreator {

	coinMarketCupClient := cli.NewCoinMarketCupClient(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", coinMarketCupApiKey,
	)
	fetcher := cli.CMCRateFetcher{
		Client: coinMarketCupClient,
	}

	storage := ports.NewInMemoryPaymentRepository()
	paymentCreateInteractor := interactors.PaymentCreator{PaymentSaver: storage, IdGenerator: ports.NewStubIdGenerator()}

	return cli.ConsolePaymentCreator{PaymentCreator: paymentCreateInteractor, RateFetcher: fetcher}
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
