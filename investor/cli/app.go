package cli

import (
	"investor/cli/asset"
	"investor/cli/payment"
	"log"
	"os"
)

type App struct {
	CreateAssetCommand   asset.ConsoleAssetCreator
	ListAssetsCommand    asset.ConsoleAssetsLister
	CreatePaymentCommand payment.ConsolePaymentCreator
	ListPaymentsCommand  payment.ConsolePaymentsLister
	CalcProfitCommand    payment.ConsoleProfitCalculator
}

func (app App) setup() CLI {
	cli := NewCLI()
	cli.AddCommand("create_asset", app.CreateAssetCommand)
	cli.AddCommand("create_payment", app.CreatePaymentCommand)
	cli.AddCommand("list_payments", app.ListPaymentsCommand)
	cli.AddCommand("list_assets", app.ListAssetsCommand)
	cli.AddCommand("calc_profit", app.CalcProfitCommand)

	return cli
}

func (app App) Run() {
	cli := app.setup()

	commands := "create_asset, list_assets, create_payment, list_payments, calc_profit"
	argsLen := len(os.Args)
	if argsLen == 2 {
		command := os.Args[1]
		cli.Run(command)
	} else if argsLen > 2 {
		log.Fatalf(
			"unexpected args, only one param expected - command with value: %s", commands,
		)
	} else {
		log.Fatalf("provide command, one of: %s", commands)
	}
}
