package payment_filters

import "time"

type Period struct {
	TimeFrom  time.Time
	TimeUntil time.Time
}
