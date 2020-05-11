package rate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Quote map[string]struct{ Price float32 `json:"price"` }

type CMCResponse struct {
	Data map[string]struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Quote  Quote  `json:"quote"`
	} `json:"data"`
}

type CoinMarketCupClient struct {
	url           string
	apiKey        string
	CurrenciesIDs map[CryptoCurrency]string
	HTTPClient    HTTPClient
}

func NewCoinMarketCupClient(url string, apiKey string) CoinMarketCupClient {
	return CoinMarketCupClient{url, apiKey,
		// todo: fetch from CoinMarketCup
		map[CryptoCurrency]string{
			BTC:  "1",
			ETH:  "1027",
			XRP:  "52",
			DASH: "131",
		},
		&http.Client{},
	}
}

func (c CoinMarketCupClient) BuildRequest(currencyID string) (*http.Request, error) {
	req, err := http.NewRequest(
		"GET", c.url, nil,
	)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("id", currencyID)
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", c.apiKey)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

type NoIDForCurrencyError struct {
	Currency CryptoCurrency
}

func (e NoIDForCurrencyError) Error() string {
	return fmt.Sprintf("Relevant id for Currency: %s, does'nt exists", e.Currency)
}

type RateRequestFailedError struct {
	Currency CryptoCurrency
	Err      string
}

func (e RateRequestFailedError) Error() string {
	return fmt.Sprintf("Failed to fetche rate for currency: %s due to error:\n%s", e.Currency, e.Err)
}

func (c CoinMarketCupClient) FetchCurrencyRate(currency CryptoCurrency) (float32, error) {
	currencyID, ok := c.CurrenciesIDs[currency]
	if !ok {
		return 0, NoIDForCurrencyError{currency}
	}
	req, err := c.BuildRequest(currencyID)
	if err != nil {
		return 0, err
	}
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return 0, RateRequestFailedError{currency, err.Error()}
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	var respData CMCResponse

	if err := json.Unmarshal(respBody, &respData); err != nil {
		return 0, err
	}
	return ParsePrice(respData, currencyID)
}

type CurrencyIDNotFoundError struct {
	CurrencyID string
}

func (e CurrencyIDNotFoundError) Error() string {
	return fmt.Sprintf("Data for currencyID: %s not found in CoinMarketCup response", e.CurrencyID)
}

type NoPriceInUSDError struct {
}

func (e NoPriceInUSDError) Error() string {
	return "USD key is not in Quote map"
}

func ParsePrice(respData CMCResponse, currencyID string) (float32, error) {
	currencyData, ok := respData.Data[currencyID]
	if !ok {
		return 0, CurrencyIDNotFoundError{currencyID}
	}

	data, ok2 := currencyData.Quote["USD"]
	if !ok2 {
		return 0, NoPriceInUSDError{}
	}
	return data.Price, nil
}
