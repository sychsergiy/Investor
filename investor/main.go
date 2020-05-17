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
	"log"
	"os"
)

func setupDependencies(coinMarketCupAPIKey string) cli.App {

	coinMarketCupClient := rate.NewCoinMarketCupClient(
		"https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", coinMarketCupAPIKey,
	)
	fetcher := rate.CMCFetcher{
		Client: coinMarketCupClient,
	}

	storage := jsonfile.NewStorage(file.NewJSON(file.NewPlainFile("storage.json")))

	assetRepo := jsonfile.NewAssetRepository(storage)
	paymentRepo := jsonfile.NewPaymentRepository(storage, assetRepo)

	paymentCreateInteractor := interactors.NewCreatePayment(paymentRepo, adapters.NewUUIDGenerator())
	paymentListInteractor := interactors.NewListPayments(paymentRepo)
	assetNamesFilterInteractor := interactors.NewPaymentAssetNamesFilter(paymentRepo)
	categoriesFilterInteractor := interactors.NewPaymentAssetCategoriesFilter(paymentRepo)
	assetCreateInteractor := interactors.NewCreateAsset(assetRepo, adapters.NewUUIDGenerator())
	assetsListInteractor := interactors.NewListAssets(assetRepo)
	calcProfitInteractor := interactors.NewCalcProfit(paymentRepo)
	calcRateFromProfitInteractor := interactors.NewCalcRateFromProfit(paymentRepo)

	paymentCreateCommand := payment.NewCreateCommand(paymentCreateInteractor, assetsListInteractor, fetcher)
	paymentsListCommand := payment.NewConsolePaymentsLister(paymentListInteractor)
	filterByAssetNamesCommand := payment.NewFilterByAssetNamesCommand(assetNamesFilterInteractor)
	filterByCategoriesCommand := payment.NewFilterByCategoriesCommand(categoriesFilterInteractor)
	calcProfitCommand := payment.NewCalcProfitCommand(calcProfitInteractor)
	calcRateFromProfitCommand := payment.NewCalcRateFromProfitCommand(
		calcRateFromProfitInteractor, assetsListInteractor,
	)

	assetCreateCommand := asset.NewCreateCommand(assetCreateInteractor)
	assetsListCommand := asset.NewListCommand(assetsListInteractor)

	return cli.App{
		CreateAssetCommand: assetCreateCommand,
		ListAssetsCommand:  assetsListCommand,

		CreatePaymentCommand:      paymentCreateCommand,
		ListPaymentsCommand:       paymentsListCommand,
		FilterByAssetNamesCommand: filterByAssetNamesCommand,
		CalcProfitCommand:         calcProfitCommand,
		CalcRateFromProfitCommand: calcRateFromProfitCommand,
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
