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
	"investor/interactors/payment_filters"
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
	assetNameFilterInteractor := payment_filters.NewAssetNameFilter(paymentRepo)
	assetCreateInteractor := interactors.NewCreateAsset(assetRepo, adapters.NewUUIDGenerator())
	assetsListInteractor := interactors.NewListAssets(assetRepo)
	calcProfitInteractor := interactors.NewCalcProfit()

	paymentCreateCommand := payment.NewConsolePaymentCreator(paymentCreateInteractor, assetsListInteractor, fetcher)
	paymentsListCommand := payment.NewConsolePaymentsLister(paymentListInteractor)
	filterByAssetNameCommand := payment.NewFilterByAssetNameCommand(assetNameFilterInteractor)
	calcProfitCommand := payment.NewConsoleProfitCalculator(paymentListInteractor, calcProfitInteractor)

	assetCreateCommand := asset.NewConsoleAssetCreator(assetCreateInteractor)
	assetsListCommand := asset.NewConsoleAssetLister(assetsListInteractor)

	return cli.App{
		CreateAssetCommand: assetCreateCommand,
		ListAssetsCommand:  assetsListCommand,

		CreatePaymentCommand:     paymentCreateCommand,
		ListPaymentsCommand:      paymentsListCommand,
		FilterByAssetNameCommand: filterByAssetNameCommand,
		CalcProfitCommand:        calcProfitCommand,
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
