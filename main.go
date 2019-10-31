package main

import (
	"Investor/asset"
	"Investor/asset/amount"
	"Investor/asset/crypto"
	"Investor/asset/payment/storage"
	"Investor/asset/period"
	"fmt"
	"time"
)

func main() {
	storage_ := storage.NewInMemory()
	asset.MakeInvestment(storage_, &amount.USD{Value: 100}, 1, time.Now())
	asset.MakeInvestment(storage_, &amount.USD{Value: 200}, 1, time.Date(2018,1,1,1,1,1,1,time.UTC))
	asset.MakeReturn(storage_, &amount.USD{Value: 50}, 0.5, time.Now())
	cryptoCurrAsset := crypto.CurrencyAsset{Storage: storage_}
	fmt.Printf("%.2f", cryptoCurrAsset.CalcProfit())
	println()
	fmt.Printf("%.2f", cryptoCurrAsset.CalcProfitOnPeriod(period.Year{Value: 2019}))
	println()
	fmt.Printf("%.2f", cryptoCurrAsset.CalcProfitOnPeriod(period.Year{Value: 2018}))
	println()

}
