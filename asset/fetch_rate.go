package asset

import (
	"Investor/asset/amount/crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type CryptoResp struct {
	Data map[string]struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Quote  map[string]struct {
			Price float32 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

func fetchCryptoCurrencyRate(c crypto.Currency, apiKey string) float32 {
	currencyID := getCoinMarketCapCurrencyID(c)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("id", currencyID)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	var respData CryptoResp

	if err := json.Unmarshal(respBody, &respData); err != nil {
		log.Fatal(err)
	}
	// todo: add key error handling
	return respData.Data[currencyID].Quote["USD"].Price
}

func getCoinMarketCapCurrencyID(c crypto.Currency) string {
	// todo:  fetch ids automatically from CoinMarketCap API
	switch c {
	case crypto.BTC:
		return "1"
	case crypto.ETH:
		return "1027"
	case crypto.XRP:
		return "52"
	case crypto.Dash:
		return "131"
	}
	panic(fmt.Sprintf("Unresolved fiat %d", c))
}

type FiatResp struct {
	ExchangeRate [] struct {
		Currency       string  `json:"currency"`
		SaleRateNB     float32 `json:"saleRateNB"`     // NB sale rate
		PurchaseRateNB float32 `json:"purchaseRateNB"` // NB purchase rate
		SaleRate       float32 `json:"saleRate"`       // optional // PrivatBank sale rate
		PurchaseRate   float32 `json:"purchaseRate"`   // optional // PrivatBank purchase rate
	} `json:"exchangeRate"`
}

func FetchUAHCurrencyRate() float32 {
	// returns UAH to USD
	// todo: add EUR and USD as base currency
	date := time.Now()
	apiUrl := "https://api.privatbank.ua/p24api/exchange_rates"

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("date", fmt.Sprintf("%d.%d.%d", date.Day()-1, date.Month(), date.Year()))
	q.Add("json", "")

	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	var fiatResp FiatResp

	if err := json.Unmarshal(respBody, &fiatResp); err != nil {
		log.Fatal(err)
	}

	for _, item := range fiatResp.ExchangeRate {
		if item.Currency == "USD" {
			return item.PurchaseRate
		}
	}
	panic("USD rate not found in response")
}
