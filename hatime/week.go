package hatime

import (
	"time"

	"github.com/jinzhu/now"
)

// caution: sunday is last day of week
func SaturdayFrom(tm time.Time) time.Time {
	now.WeekStartDay = time.Monday
	return now.With(tm).Sunday().Add(-time.Hour * 24)
}

// caution: sunday is last day of week
func SundayFrom(tm time.Time) time.Time {
	now.WeekStartDay = time.Monday
	return now.With(tm).Sunday()
}

// caution: sunday is last day of week
func SaturdayOf(tm time.Time) time.Time {
	now.WeekStartDay = time.Monday
	return now.With(tm).EndOfWeek().Add(-time.Hour*24 + 1)
}

// caution: sunday is last day of week
func SundayOf(tm time.Time) time.Time {
	now.WeekStartDay = time.Monday
	return now.With(tm).EndOfWeek().Add(1)
}
