package hakit_test

import (
	"testing"

	"github.com/johnhaha/hakit/hatest"
	"github.com/johnhaha/hakit/havalid"
)

func TestIP(t *testing.T) {
	checks := []hatest.InOut{
		{
			Input:  "not ip",
			Output: 0,
		},
		{
			Input:  "10.10.10.10",
			Output: 4,
		},
		{
			Input:  "2001:db8::ff00:42:8329",
			Output: 6,
		},
	}
	for _, c := range checks {
		res := havalid.IPCheck(c.Input.(string))
		if res != c.Output.(int) {
			t.Fatal(res)
		}
	}
}
