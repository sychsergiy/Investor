package crypto

import (
	"Investor/asset"
	"Investor/asset/amount"
	"Investor/asset/payment/storage"
	"Investor/asset/period"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCurrencyAsset_CalcProfitOnPeriod(t *testing.T) {
	storage_ := storage.NewInMemory()
	asset.MakeInvestment(storage_, &amount.USD{Value: 100}, 1, time.Now())
	asset.MakeInvestment(storage_, &amount.USD{Value: 200}, 1, time.Date(
		2018, 1, 1, 1, 1, 1, 1, time.UTC))
	asset.MakeReturn(storage_, &amount.USD{Value: 50}, 0.5, time.Now())
	cryptoCurrAsset := CurrencyAsset{Storage: storage_}

	assert.Equal(t, "0.67", fmt.Sprintf("%.2f", cryptoCurrAsset.CalcProfit()))
	assert.Equal(t, "1.00", fmt.Sprintf("%.2f", cryptoCurrAsset.CalcProfitOnPeriod(period.Year{Value: 2019})))
	assert.Panics(t, func() { cryptoCurrAsset.CalcProfitOnPeriod(period.Year{Value: 2018}) }) // no investments during 2018
}

func TestCurrencyAsset_CalcCurrencyRate(t *testing.T) {
	storage_ := storage.NewInMemory()
	date := time.Date(2019, 1, 1, 1, 1, 1, 1, time.UTC)
	asset.MakeInvestment(storage_, &amount.USD{Value: 100}, 1, date)
	asset.MakeInvestment(storage_, &amount.USD{Value: 200}, 1, date)
	asset.MakeReturn(storage_, &amount.USD{Value: 100}, 0.5, date)
	asset.MakeReturn(storage_, &amount.USD{Value: 50}, 0.5, date)

	currencyAsset := CurrencyAsset{storage_}
	currencyRate := currencyAsset.CalcCurrencyRate(1.5)

	assert.Equal(t, float32(300), currencyRate)
}

func TestCurrencyAsset_CalcCurrencyRateOnPeriod(t *testing.T) {
	storage_ := storage.NewInMemory()
	dateIn := time.Date(2019, 1, 1, 1, 1, 1, 1, time.UTC)
	dateOut := time.Date(2018, 1, 1, 1, 1, 1, 1, time.UTC)
	asset.MakeInvestment(storage_, &amount.USD{Value: 100}, 1, dateIn)
	asset.MakeInvestment(storage_, &amount.USD{Value: 200}, 1, dateIn)
	asset.MakeReturn(storage_, &amount.USD{Value: 100}, 0.5, dateIn)
	asset.MakeReturn(storage_, &amount.USD{Value: 50}, 0.5, dateIn)

	asset.MakeInvestment(storage_, &amount.USD{Value: 30}, 1, dateOut)
	asset.MakeReturn(storage_, &amount.USD{Value: 30}, 0.5, dateOut)

	currencyAsset := CurrencyAsset{storage_}
	currencyRate := currencyAsset.CalcCurrencyRateOnPeriod(1.5, period.Year{Value: 2019})

	assert.Equal(t, float32(300), currencyRate)
}
