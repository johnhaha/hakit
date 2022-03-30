package hakit_test

import (
	"testing"
	"time"

	"github.com/johnhaha/hakit/hatime"
)

func TestGetTimeIn(t *testing.T) {
	chicago, _ := time.LoadLocation("America/Chicago")
	now := time.Now().In(chicago)
	tm := hatime.GetTodayBeginIn(now)
	t.Fatal(tm)
}

func TestGetDay(t *testing.T) {
	d := hatime.SundayOf(time.Now().Add(time.Hour * 24 * 5))
	t.Fatal(d)
}
