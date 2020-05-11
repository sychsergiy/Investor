package rate

import (
	"investor/entities/asset"
)

type CMCFetcher struct {
	Client        CoinMarketCupClient
	CurrenciesIDs map[CryptoCurrency]string
}

func (f CMCFetcher) Fetch(a asset.Asset) (Rate, error) {
	rate, err := f.Client.FetchCurrencyRate(CryptoCurrency(a.Name()))
	if err != nil {
		return 0, err
	}

	return Rate(rate), nil
}
