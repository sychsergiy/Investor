package main

import (
	"fmt"
	"investor/asset/crypto_currency"
	"log"
	"os"
)

func setupDependencies(coinMarketCupApiKey string) crypto_currency.RateFetcher {
	coinMarketCupClient := crypto_currency.NewCoinMarketCupClient(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", coinMarketCupApiKey,
	)
	fetcher := crypto_currency.RateFetcher{
		Client: coinMarketCupClient,
	}
	return fetcher
}
func main() {
	apiKey := os.Getenv("COIN_MARKET_CAP_API_KEY")
	if apiKey == "" {
		log.Fatal("COIN_MARKET_CAP_API_KEY env var not provided")
	}
	fetcher := setupDependencies(apiKey)
	rate, err := fetcher.Fetch(crypto_currency.NewAsset(crypto_currency.BTC))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("BTC rate: %f", rate)
}
