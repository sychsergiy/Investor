package payment

import "time"

type Period interface {
	From() time.Time
	Until() time.Time
}

type MonthPeriod struct {
	year  int
	month time.Month
}

func (p MonthPeriod) From() time.Time {
	return createDate(p.year, p.month)
}

func createDate(year int, month time.Month) time.Time {
	return time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
}

func (p MonthPeriod) Until() time.Time {
	return createDate(p.year, p.month+1)
}

func NewMonthPeriod(year int, month time.Month) MonthPeriod {
	return MonthPeriod{year, month}
}

type YearPeriod struct {
	year int
}

func (p YearPeriod) From() time.Time {
	return createDate(p.year, time.January)
}

func (p YearPeriod) Until() time.Time {
	return createDate(p.year+1, time.January)
}

func NewYearPeriod(year int) YearPeriod {
	return YearPeriod{year}
}

type DurationPeriod struct {
	from  time.Time
	until time.Time
}

func (p DurationPeriod) From() time.Time {
	return p.from
}

func (p DurationPeriod) Until() time.Time {
	return p.until
}

func NewDurationPeriod(from time.Time, until time.Time) DurationPeriod {
	return DurationPeriod{from: from, until: until}
}
