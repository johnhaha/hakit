package hatime

import (
	"time"

	"github.com/jinzhu/now"
)

func SaturdayFrom(tm time.Time) time.Time {
	return now.With(tm).Sunday().Add(-time.Hour * 24)
}

func SundayFrom(tm time.Time) time.Time {
	return now.With(tm).Sunday()
}

func SaturdayOf(tm time.Time) time.Time {
	return now.EndOfWeek().Add(-time.Hour*24 + 1)
}

// caution: sunday is last day of week
func SundayOf(tm time.Time) time.Time {
	return now.EndOfWeek().Add(1)
}
