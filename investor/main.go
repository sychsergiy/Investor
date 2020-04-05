package main

import (
	"investor/asset/crypto_currency"
	"investor/interactors"
	"investor/payment"
	"log"
	"os"
)

func setupDependencies(coinMarketCupApiKey string) interactors.PaymentCreator {

	coinMarketCupClient := crypto_currency.NewCoinMarketCupClient(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", coinMarketCupApiKey,
	)
	fetcher := crypto_currency.RateFetcher{
		Client: coinMarketCupClient,
	}

	storage := payment.NewInMemoryStorage()
	return interactors.PaymentCreator{Storage: storage, RateFetcher: fetcher}
}
func main() {
	apiKey := os.Getenv("COIN_MARKET_CAP_API_KEY")
	if apiKey == "" {
		log.Fatal("COIN_MARKET_CAP_API_KEY env var not provided")
	}
	paymentCreator := setupDependencies(apiKey)
	paymentCreator.Create()
}
