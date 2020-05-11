package payment

import (
	"investor/entities/asset"
	"time"
)

func CreatePaymentWithAmount(paymentType Type, amount, assetAmount float32) Payment {
	a := asset.NewPlainAsset("gold", asset.PreciousMetal, "gold")
	date := time.Date(2019, 30, 12, 11, 58, 0, 0, time.UTC)
	return NewPlainPayment("test", assetAmount, amount, a, date, paymentType)
}

func CreatePayment(id string, year int) Payment {
	testAsset := asset.NewPlainAsset("test", asset.CryptoCurrency, "test")
	creationTime := CreateYearDate(year)
	return NewPlainPayment(id, 0, 0, testAsset, creationTime, Invest)
}

func CreatePaymentWithAsset(id, assetID string, year int) Payment {
	testAsset := asset.NewPlainAsset(assetID, asset.CryptoCurrency, "test")
	creationTime := CreateYearDate(year)
	return NewPlainPayment(id, 0, 0, testAsset, creationTime, Invest)
}

func CreateYearDate(year int) time.Time {
	return CreateMonthDate(year, 0)
}

func CreateMonthDate(year int, month time.Month) time.Time {
	return CreateDayDate(year, month, 0)
}

func CreateDayDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

type PeriodMock struct {
	TimeFrom  time.Time
	TimeUntil time.Time
}

func (p PeriodMock) From() time.Time {
	return p.TimeFrom
}

func (p PeriodMock) Until() time.Time {
	return p.TimeUntil
}

type ProxyMock struct {
	Payment
	assetFunc func() (asset.Asset, error)
}

func (ppm ProxyMock) Asset() (asset.Asset, error) {
	return ppm.assetFunc()
}

func NewProxyMock(p Payment, assetFunc func() (asset.Asset, error)) ProxyMock {
	return ProxyMock{p, assetFunc}
}
