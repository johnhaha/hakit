package hatime

import (
	"time"

	"github.com/jinzhu/now"
)

func GetTodayBegin() time.Time {
	tm := now.BeginningOfDay()
	return tm
}

func GetTodayEnd() time.Time {
	tm := now.EndOfDay()
	return tm
}

func GetTodayBeginIn(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func GetTodayEndIn(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}
