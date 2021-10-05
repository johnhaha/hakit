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
