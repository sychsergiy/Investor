package cli

import (
	"investor/cli/asset"
	"investor/cli/payment"
	"log"
	"os"
)

type App struct {
	CreateAssetCommand        asset.CreateAssetCommand
	ListAssetsCommand         asset.ListAssetsCommand
	CreatePaymentCommand      payment.CreatePaymentCommand
	ListPaymentsCommand       payment.ListPaymentsCommand
	CalcProfitCommand         payment.CalcProfitCommand
	FilterByAssetNameCommand  payment.FilterByAssetNameCommand
	FilterByCategoriesCommand payment.FilterByCategoriesCommand

	cli *CLI
}

func (app *App) setup() {
	app.cli = NewCLI()
	app.cli.AddCommand("create_asset", app.CreateAssetCommand)
	app.cli.AddCommand("create_payment", app.CreatePaymentCommand)
	app.cli.AddCommand("list_payments", app.ListPaymentsCommand)
	app.cli.AddCommand("list_assets", app.ListAssetsCommand)
	app.cli.AddCommand("calc_profit", app.CalcProfitCommand)

	app.cli.AddCommand("filter_by_asset_name", app.FilterByAssetNameCommand)
	app.cli.AddCommand("filter_by_categories", app.FilterByCategoriesCommand)
}

func (app App) Run() {
	app.setup()

	commands := app.cli.AvailableCommands()
	argsLen := len(os.Args)
	if argsLen == 2 {
		command := os.Args[1]
		app.cli.Run(command)
	} else if argsLen > 2 {
		log.Fatalf(
			"unexpected args, only one param expected - command with value: %s", commands,
		)
	} else {
		log.Fatalf("provide command, one of: %s", commands)
	}
}
