package cli

import (
	"investor/cli/asset"
	"investor/cli/payment"
	"log"
	"os"
)

type App struct {
	CreateAssetCommand   asset.ConsoleAssetCreator
	CreatePaymentCommand payment.ConsolePaymentCreator
	ListPaymentsCommand  payment.ConsolePaymentsLister
}

func (app App) setup() CLI {
	cli := NewCLI()
	cli.AddCommand("create_asset", app.CreateAssetCommand)
	cli.AddCommand("create_payment", app.CreatePaymentCommand)
	cli.AddCommand("list_payments", app.ListPaymentsCommand)

	return cli
}

func (app App) Run() {
	cli := app.setup()

	commands := "create_asset, create_payment, list_payments"
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
