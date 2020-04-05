package period

import (
	"time"
)

type Period interface {
	Contains(date time.Time) bool
}

type Month struct {
	Year  int
	Month time.Month
}

func (m Month) Contains(date time.Time) bool {
	from := newDate(m.Year, m.Month, 0)
	to := newDate(m.Year, m.Month+1, 0)
	return Duration{from, to}.Contains(date)
}

type Year struct {
	Value int
}

func (y Year) Contains(date time.Time) bool {
	from := newDate(y.Value, 1, 1)
	to := newDate(y.Value+1, 1, 1)
	return Duration{from, to}.Contains(date)
}

type Duration struct {
	From time.Time
	To   time.Time
}

func (d Duration) Contains(date time.Time) bool {
	return d.To.After(date) && d.From.Before(date)
}

func newDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
