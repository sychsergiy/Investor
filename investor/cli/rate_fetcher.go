package cli

import (
	"investor/entities/asset"
	"investor/entities/asset/crypto_currency"
)

type CMCRateFetcher struct {
	Client        CoinMarketCupClient
	CurrenciesIDs map[crypto_currency.CryptoCurrency]string
}

func (f CMCRateFetcher) Fetch(a asset.Asset) (Rate, error) {
	rate, err := f.Client.FetchCurrencyRate(crypto_currency.CryptoCurrency(a.Name))
	if err != nil {
		return 0, err
	}

	return Rate(rate), nil
}
