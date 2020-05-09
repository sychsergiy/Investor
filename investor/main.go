package main

import (
	"investor/adapters"
	"investor/adapters/repositories/jsonfile"
	"investor/cli"
	"investor/cli/asset"
	"investor/cli/payment"
	"investor/cli/payment/rate_fetcher"
	"investor/helpers/file"
	"investor/interactors"
	"log"
	"os"
)

func setupDependencies(coinMarketCupApiKey string) cli.App {

	coinMarketCupClient := rate_fetcher.NewCoinMarketCupClient(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", coinMarketCupApiKey,
	)
	fetcher := rate_fetcher.CMCRateFetcher{
		Client: coinMarketCupClient,
	}

	storage := jsonfile.NewStorage(file.NewJsonFile(file.NewPlainFile("storage.json")))

	assetRepo := jsonfile.NewAssetRepository(storage)
	paymentRepo := jsonfile.NewPaymentRepository(storage, assetRepo)

	paymentCreateInteractor := interactors.CreatePayment{Repository: paymentRepo, IdGenerator: adapters.NewUUIDGenerator()}
	paymentListInteractor := interactors.ListPayments{Repository: paymentRepo}
	assetCreateInteractor := interactors.NewCreateAsset(assetRepo, adapters.NewUUIDGenerator())

	paymentCreateCommand := payment.ConsolePaymentCreator{PaymentCreator: paymentCreateInteractor, RateFetcher: fetcher}
	paymentsListCommand := payment.NewConsolePaymentsLister(paymentListInteractor)
	assetCreateCommand := asset.NewConsoleAssetCreator(assetCreateInteractor)

	return cli.App{
		CreateAssetCommand:   assetCreateCommand,
		CreatePaymentCommand: paymentCreateCommand,
		ListPaymentsCommand:  paymentsListCommand,
	}
}
func main() {
	apiKey := os.Getenv("COIN_MARKET_CAP_API_KEY")
	if apiKey == "" {
		log.Fatal("COIN_MARKET_CAP_API_KEY env var not provided")
	}
	app := setupDependencies(apiKey)
	app.Run()
}
