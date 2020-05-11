package rate_fetcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"investor/entities/asset/crypto_currency"
	"io/ioutil"
	"net/http"
	"testing"
)

func mockRateResponse(currencyID string, quote Quote) RateResponse {
	return RateResponse{
		Data: map[string]struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
			Quote  Quote  `json:"quote"`
		}{currencyID: {1, "test", "test", quote},},
	}
}

func TestParsePrice(t *testing.T) {
	price := float32(10)
	currencyID := "BTC"
	data := mockRateResponse(currencyID, Quote{"USD": {price}})
	result, err := ParsePrice(data, currencyID)

	if result != 10 {
		t.Errorf("With valid rate data expected price: %f", price)
	} else {
		if err != nil {
			t.Errorf("Expected no error(nil) when price value returned.")
		}
	}
}

func TestParsePriceWrongCurrencyID(t *testing.T) {
	data := mockRateResponse("BTC", Quote{"USD": {Price: 10}})
	result, err := ParsePrice(data, "NotExistent")

	expectedErr := CurrencyIDNotFoundError{"NotExistent"}
	if err != expectedErr {
		t.Errorf("Not expected error when Currency id is not found")
	} else {
		if result != 0 {
			t.Errorf("Expected 0 value when error returned.")
		}
	}
}

func TestParsePriceNoUSDInQuote(t *testing.T) {
	data := mockRateResponse("BTC", Quote{"NotExistent": {Price: 10}})
	result, err := ParsePrice(data, "BTC")

	expectedErr := NoPriceInUSDError{}
	if err != expectedErr {
		t.Errorf("Not expected error when there is no USD key in Quote map.")
	} else {
		if result != 0 {
			t.Errorf("Expected 0 value when error returned.")
		}
	}
}

func TestCoinMarketCupClient_BuildRequest(t *testing.T) {
	url := "http://url.com"
	client := NewCoinMarketCupClient(url, "api_key")
	req, err := client.BuildRequest("1")
	if err != nil {
		t.Errorf("Not expected error during request build: %s", err)
	}

	accepts := req.Header.Get("Accepts")
	if accepts != "application/json" {
		t.Errorf("Expected value for 'Accespts' header: %s, got: %s", "applicaiton/json", accepts)
	}

	apiKey := req.Header.Get("X-CMC_PRO_API_KEY")
	if apiKey != "api_key" {
		t.Errorf("Expected value for 'X-CMC_PRO_API_KEY' header: %s, got: %s", "X-CMC_PRO_API_KEY", accepts)
	}

	if req.URL.RawQuery != "id=1" {
		t.Errorf("Not expected value for raw query")
	}
}

type HttpClientErrorMock struct {
}

func (c HttpClientErrorMock) Do(req *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("mocked error")
}

type HttpClientMock struct{}

func (c HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	rateData := mockRateResponse("1", Quote{"USD": {10}})
	marshalled, _ := json.Marshal(rateData)
	body := ioutil.NopCloser(bytes.NewReader(marshalled))
	return &http.Response{
		StatusCode: 200,
		Body:       body,
	}, nil
}

func TestCoinMarketCupClient_FetchCurrencyRate_WrongCurrencyID(t *testing.T) {
	client := NewCoinMarketCupClient("http://url.com", "api_key")
	client.HttpClient = HttpClientMock{}
	expectedErr := NoIDForCurrencyError{"wrong"}

	rate, err := client.FetchCurrencyRate("wrong")
	if err != expectedErr {
		t.Errorf("Currency rate id not found error expected.")
		if rate != 0 {
			t.Errorf("Zero rate value expected when error returned.")
		}
	}

}

func TestCoinMarketCupClient_FetchCurrencyRate_RequestError(t *testing.T) {
	client := NewCoinMarketCupClient("http://url.com", "api_key")
	client.HttpClient = HttpClientErrorMock{}
	expectedErr := RateRequestFailedError{crypto_currency.BTC, "mocked error"}

	rate, err := client.FetchCurrencyRate(crypto_currency.BTC)
	if err != expectedErr {
		t.Errorf("Rate request failed error expected.")
	}
	if rate != 0 {
		t.Errorf("Zero rate value expected when error returned.")
	}
}

func TestCoinMarketCupClient_FetchCurrencyRate_Ok(t *testing.T) {
	client := NewCoinMarketCupClient("http://url.com", "api_key")
	client.HttpClient = HttpClientMock{}
	rate, err := client.FetchCurrencyRate(crypto_currency.BTC)
	if err != nil {
		t.Errorf("No error expected.")
	}
	if rate != 10 {
		t.Errorf("Expected price value: 10, got %f", rate)
	}
}
