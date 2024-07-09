package hatime

import (
	"time"

	"github.com/jinzhu/now"
)

func StartHourOf(tm time.Time) time.Time {
	return now.With(tm).BeginningOfHour()
}

func EndHourOf(tm time.Time) time.Time {
	return now.With(tm).EndOfHour()
}

func NextHourOf(tm time.Time, hour int) time.Time {
	if hour <= tm.Hour() {
		tm = tm.Add(time.Hour * 24)
	}
	return time.Date(tm.Year(), tm.Month(), tm.Day(), hour, 0, 0, 0, tm.Location())
}
