package payment

import (
	"testing"
	"time"
)

func verifyPeriod(t *testing.T, p Period, expectedFrom, expectedUntil time.Time) {
	gotFrom := p.From()
	if gotFrom != expectedFrom {
		t.Errorf("Expected from value: %s, but got: %s", expectedFrom, gotFrom)
	}
	gotUntil := p.Until()
	if gotUntil != expectedUntil {
		t.Errorf("Expected until value: %s, but got: %s", expectedUntil, gotUntil)
	}
}

func TestMonthPeriod(t *testing.T) {
	from := time.Date(2020, 2, 0, 0, 0, 0, 0, time.UTC)
	until := time.Date(2020, 3, 0, 0, 0, 0, 0, time.UTC)

	p := NewMonthPeriod(2020, time.February)
	verifyPeriod(t, p, from, until)
}

func TestYearPeriod(t *testing.T) {
	from := time.Date(2019, time.December, 31, 0, 0, 0, 0, time.UTC)
	until := time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC)

	p := NewYearPeriod(2020)
	verifyPeriod(t, p, from, until)
}
