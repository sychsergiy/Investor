package crypto_currency

import (
	"investor/asset"
)

type RateFetcher struct {
	Client        CoinMarketCupClient
	CurrenciesIDs map[CryptoCurrency]string
}

func (f RateFetcher) Fetch(a asset.Asset) (asset.Rate, error) {
	rate, err := f.Client.FetchCurrencyRate(CryptoCurrency(a.Name))
	if err != nil {
		return 0, err
	}

	return asset.Rate(rate), nil
}
