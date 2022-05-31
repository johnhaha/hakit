package hatime

import (
	"time"

	"github.com/jinzhu/now"
)

func StartHourOf(time time.Time) time.Time {
	return now.With(time).BeginningOfHour()
}

func EndHourOf(time time.Time) time.Time {
	return now.With(time).EndOfHour()
}
