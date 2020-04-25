package cli

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	timeStr := "2019-12-30 11:58:59"
	date, err := ParseTime(timeStr)

	if err != nil {
		t.Errorf("Failed to parse time due to err: %s", err)
		if date != time.Date(2019, 30, 12, 11, 58, 59, 0, time.UTC) {
			t.Errorf("Wrong parsed date value")
		}
	}
}
