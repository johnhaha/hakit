package hakit_test

import (
	"strconv"
	"testing"

	"github.com/johnhaha/hakit/hadata"
	"github.com/johnhaha/hakit/hatest"
)

func TestUpVersion(t *testing.T) {
	v := []hatest.StringInOut{
		{
			Input:  "1.0.1",
			Output: "2.0.0",
		},
		{
			Input:  "1.1.1",
			Output: "2.0.0",
		},
		{
			Input:  "v1.0.1",
			Output: "v2.0.0",
		},
	}
	for i, x := range v {
		res, err := hadata.UpgradeVersion(x.Input, 1)
		if err != nil {
			t.Fatal(err)
		}
		if res != x.Output {
			t.Fatal(strconv.Itoa(i), res)
		}
	}
}

func TestIsIn(t *testing.T) {
	vs := []hatest.InOut{
		{
			Input:  "ok",
			Output: true,
		},
		{
			Input:  "not ok",
			Output: false,
		},
	}
	for _, v := range vs {
		res := hadata.IsInStringSlice([]string{"ok", "yes"}, v.Input.(string))
		if res != v.Output {
			t.Fatal("you fucked")
		}
	}
}

func TestSwapSlice(t *testing.T) {
	ios := []hatest.InOut{
		{
			Input:  []string{"1", "2"},
			Output: []string{"2", "1"},
		},
		{
			Input:  []string{"x", "y"},
			Output: []string{"y", "x"},
		},
	}
	for _, io := range ios {
		res, err := hadata.SwapStringSlice(io.Input.([]string))
		if err != nil {
			t.Fatal(err)
		}
		if res[0] != io.Output.([]string)[0] {
			t.Fatal("oh no")
		}
	}
}
