package main

import (
	"investor/cli"
	"investor/interactors"
	"investor/ports"
	"log"
	"os"
)

func setupDependencies(coinMarketCupApiKey string) interactors.PaymentCreator {

	coinMarketCupClient := cli.NewCoinMarketCupClient(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", coinMarketCupApiKey,
	)
	fetcher := cli.CMCRateFetcher{
		Client: coinMarketCupClient,
	}

	storage := ports.NewInMemoryStorage()
	return interactors.PaymentCreator{Storage: storage, RateFetcher: fetcher, IdGenerator: ports.NewStubIdGenerator()}
}
func main() {
	apiKey := os.Getenv("COIN_MARKET_CAP_API_KEY")
	if apiKey == "" {
		log.Fatal("COIN_MARKET_CAP_API_KEY env var not provided")
	}
	paymentCreator := setupDependencies(apiKey)
	paymentCreator.CreateFromCLI()
}
