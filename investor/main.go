package main

import (
	"investor/adapters"
	"investor/adapters/repositories/jsonfile"
	"investor/cli"
	"investor/cli/asset"
	"investor/cli/payment"
	"investor/cli/payment/rate"
	"investor/helpers/file"
	"investor/interactors"
	"investor/interactors/payment_filters"
	"log"
	"os"
)

func setupDependencies(coinMarketCupApiKey string) cli.App {

	coinMarketCupClient := rate.NewCoinMarketCupClient(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", coinMarketCupApiKey,
	)
	fetcher := rate.CMCFetcher{
		Client: coinMarketCupClient,
	}

	storage := jsonfile.NewStorage(file.NewJSONFile(file.NewPlainFile("storage.json")))

	assetRepo := jsonfile.NewAssetRepository(storage)
	paymentRepo := jsonfile.NewPaymentRepository(storage, assetRepo)

	paymentCreateInteractor := interactors.CreatePayment{Repository: paymentRepo, IDGenerator: adapters.NewUUIDGenerator()}
	paymentListInteractor := interactors.ListPayments{Repository: paymentRepo}
	assetNamesFilterInteractor := payment_filters.NewAssetNamesFilter(paymentRepo)
	categoriesFilterInteractor := payment_filters.NewAssetCategoriesFilter(paymentRepo)
	assetCreateInteractor := interactors.NewCreateAsset(assetRepo, adapters.NewUUIDGenerator())
	assetsListInteractor := interactors.NewListAssets(assetRepo)
	calcProfitInteractor := interactors.NewCalcProfit(paymentRepo)

	paymentCreateCommand := payment.NewCreatePaymentCommand(paymentCreateInteractor, assetsListInteractor, fetcher)
	paymentsListCommand := payment.NewConsolePaymentsLister(paymentListInteractor)
	filterByAssetNamesCommand := payment.NewFilterByAssetNamesCommand(assetNamesFilterInteractor)
	filterByCategoriesCommand := payment.NewFilterByCategoriesCommand(categoriesFilterInteractor)
	calcProfitCommand := payment.NewCalcProfitCommand(calcProfitInteractor)

	assetCreateCommand := asset.NewCreateAssetCommand(assetCreateInteractor)
	assetsListCommand := asset.NewListAssetsCommand(assetsListInteractor)

	return cli.App{
		CreateAssetCommand: assetCreateCommand,
		ListAssetsCommand:  assetsListCommand,

		CreatePaymentCommand:      paymentCreateCommand,
		ListPaymentsCommand:       paymentsListCommand,
		FilterByAssetNamesCommand: filterByAssetNamesCommand,
		CalcProfitCommand:         calcProfitCommand,
		FilterByCategoriesCommand: filterByCategoriesCommand,
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
