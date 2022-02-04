package hakit_test

import (
	"strconv"
	"testing"

	"github.com/johnhaha/hakit/hadata"
	"github.com/johnhaha/hakit/hatest"
)

func TestUpVersion(t *testing.T) {
	v := []hatest.InOut[string, string]{
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
	vs := []hatest.InOut[string, bool]{
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
		res := hadata.IsInSlice([]string{"ok", "yes"}, v.Input)
		if res != v.Output {
			t.Fatal("you fucked")
		}
	}
}

func TestSwapSlice(t *testing.T) {
	ios := []hatest.InOut[[]string, []string]{
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
		ipt := io.Input
		err := hadata.SwapSlice(&ipt)
		if err != nil {
			t.Fatal(err)
		}
		if ipt[0] != io.Output[0] {
			t.Fatal("oh no")
		}
	}
}

func TestStringBuilder(t *testing.T) {
	data := hadata.NewStringBinder().BindString("ok", " ok").BindInt(1, 2, 3).Value()
	if data != "ok ok123" {
		t.Fatal(data)
	}
}

func TestVersionToInt(t *testing.T) {
	res, err := hadata.VersionToInt("V0.2.4", 3)
	if err != nil {
		t.Fatal(err)
	}
	if res != 1000002004 {
		t.Fatal(res)
	}
}

func TestVersionCleaner(t *testing.T) {
	res := hadata.VersionCleaner("V111vvvvvv")
	if res != "111" {
		t.Fatal(res)
	}
}

func TestBinder(t *testing.T) {
	var b hadata.Binder = "binder"
	res := b.With("o", "k")
	if res != "binderok" {
		t.Fatal(res)
	}
}

func TestGenerteRandomString(t *testing.T) {
	res := hadata.GenerateRandomString(5)
	if len(res) != 5 {
		t.Fatal(res)
	}
	// t.Fatal(res)
}

func TestGetName(t *testing.T) {
	type Request struct {
	}
	data := new(Request)
	name := hadata.GetStructName(data)
	if name != "Request" {
		t.Fatal(name)
	}
}

func TestGetPointerData(t *testing.T) {
	i := new(int)
	*i = 9
	d, err := hadata.GetPointerData(i)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(d)
}

func TestReadStructTagData(t *testing.T) {
	type Sample struct {
		T1 string `json:"t1" test:"1"`
		T2 string `json:"t2" test:"2"`
		T3 string `json:"t3"`
	}
	data := Sample{
		T1: "1",
		T2: "",
		T3: "3",
	}
	res, err := hadata.ReadStructTagData(data, "test", "t2")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := res["t1"]; !ok {
		t.Fatal(res)
	}
	t.Fatal(res)
}

func TestReadTagMark(t *testing.T) {
	type Sample struct {
		T1 string `json:"t1" test:"1"`
		T2 string `json:"t2" test:"2"`
		T3 string `json:"t3"`
	}
	data := Sample{
		T1: "1",
		T2: "",
		T3: "3",
	}
	k, res, err := hadata.LookUpFirstTagMark(data, "test", "1")
	if err != nil {
		t.Fatal(err)
	}

	if k != "t1" {
		t.Fatal(k)
	}
	if res != "1" {
		t.Fatal(res)
	}
}

func TestSliceCheck(t *testing.T) {
	test1 := []string{"s", "2"}
	test2 := "x"
	check1, _ := hadata.CheckInterfaceDataIsSlice(test1)
	check2, _ := hadata.CheckInterfaceDataIsSlice(test2)
	if !check1 {
		t.Fatal()
	}
	if check2 {
		t.Fatal()
	}
}

func TestGetSliceFromInterface(t *testing.T) {
	test1 := []string{"s", "2"}
	test2 := &test1
	res, err := hadata.GetSliceFromInterface(test2)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 2 {
		t.Fatal()
	}
}
