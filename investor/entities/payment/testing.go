package payment

import (
	"investor/entities/asset"
	"time"
)

func CreatePaymentWithAmount(type_ Type, amount, assetAmount float32) Payment {
	a := asset.NewPlainAsset("gold", asset.PreciousMetal, "gold")
	date := time.Date(2019, 30, 12, 11, 58, 0, 0, time.UTC)
	return NewPlainPayment("test", assetAmount, amount, a, date, type_)
}

func CreatePayment(id string, year int) Payment {
	testAsset := asset.NewPlainAsset("test", asset.CryptoCurrency, "test")
	creationTime := CreateYearDate(year)
	return NewPlainPayment(id, 0, 0, testAsset, creationTime, Invest)
}

func CreatePaymentWithAsset(id, assetId string, year int) Payment {
	testAsset := asset.NewPlainAsset(assetId, asset.CryptoCurrency, "test")
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

type PaymentProxyMock struct {
	Payment
	assetFunc func() (asset.Asset, error)
}

func (ppm PaymentProxyMock) Asset() (asset.Asset, error) {
	return ppm.assetFunc()
}

func NewPaymentProxyMock(p Payment, assetFunc func() (asset.Asset, error)) PaymentProxyMock {
	return PaymentProxyMock{p, assetFunc}
}
