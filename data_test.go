package hakit_test

import (
	"strconv"
	"testing"

	"github.com/johnhaha/hakit/hadata"
	"github.com/johnhaha/hakit/hatest"
)

func BenchmarkXxx(b *testing.B) {
	data := []string{"a", "b", "c", "c", "c", "c", "c", "c", "c", "a", "sd"}
	for i := 0; i < b.N; i++ {
		hadata.RemWhereX(data, func(x string) bool {
			return x == "c"
		})
	}
}

func TestRemWhere(t *testing.T) {
	data := []string{"a", "b", "c", "c", "c", "c", "c", "c", "c", "a", "sd"}
	res := hadata.RemWhereX(data, func(x string) bool {
		return x == "c"
	})
	t.Fatal(res)
}

func TestTitleWord(t *testing.T) {
	word := "test"
	res := hadata.TitleWord(word)
	if res != "Test" {
		t.Fatal(res)
	}
}

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
	res, err := hadata.VersionToInt("v1.2.104", 3)
	if err != nil {
		t.Fatal(err)
	}
	if res != 1001002104 {
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

func TestGetNameInLowerCase(t *testing.T) {
	type Request struct {
	}
	data := new(Request)
	name := hadata.GetStructNameInLowerCase(data)
	if name != "request" {
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

func TestFindTagFiled(t *testing.T) {
	type SampleData struct {
		T1 string `json:"t1" hamo:"index"`
		T2 string `json:"t2" hamo:"location"`
	}
	res := hadata.FindTagFiled[SampleData]("hamo", "location")
	if res[0] != "t2" {
		t.Fatal(res)
	}
}

func TestFindTypeFiled(t *testing.T) {
	type SampleData struct {
		T1 string `json:"t1" hamo:"index"`
		T2 string `json:"t2" hamo:"location"`
	}
	res := hadata.FindTypeFiled[SampleData, string]()
	if res[0] != "t1" {
		t.Fatal(res)
	}
}

func TestBinderLine(t *testing.T) {
	x := hadata.NewStringBinder()
	x.BindWithNewLine("xxx", "yyy", "xxx")
	// t.Fatal(x.Value())
}

func FuzzGetRandom(f *testing.F) {
	for i := 0; i < 100; i++ {
		f.Add(i, i+hadata.GetRandomNumber(1, 100))
	}
	f.Fuzz(func(t *testing.T, min int, max int) {
		if max == min {
			t.Fatal(max)
		}
		r := hadata.GetRandomNumber(min, max)
		if r < min || r >= max {
			t.Fatal(r)
		}
	})
}

func TestRandom(t *testing.T) {
	r := hadata.GetManyRandomNumber(0, 10, 3)
	t.Fatal(r)
}

func TestAdder(t *testing.T) {
	adder := hadata.NewAdder(3, 10, 1)
	res := adder.Add(0)
	if res != 3 {
		t.Fatal(res)
	}
	res = adder.Add(5)
	if res != 6 {
		t.Fatal(res)
	}
	res = adder.Add(10)
	if res != 10 {
		t.Fatal(res)
	}
}

func TestFloatCompare(t *testing.T) {
	res := hadata.FloatCompare(1.22222, 1.21433, 2)
	t.Fatal(res)
}

func TestSplitStringOnUpperCase(t *testing.T) {
	data := "My"
	res := hadata.SplitStringOnUpperCase(data)
	t.Fatal(res)
}

func TestCamelToSnake(t *testing.T) {
	data := "John"
	res := hadata.CaseCamelToSnake(data)
	t.Fatal(res)
}

func TestMapFromSlice(t *testing.T) {
	ot := hadata.MapFromSlice([]string{"w", "s"}, func(t string) string {
		return t
	})
	t.Fatal(ot)
}

func TestMap(t *testing.T) {
	data := []string{"1", "22"}
	res := hadata.Map(data, func(t string) int {
		return len(t)
	})
	t.Fatal(res)
}

func TestFold(t *testing.T) {
	data := []int{1, 2, 3}
	res := hadata.Fold(data, 0, func(v int, t int) int {
		return v + t
	})
	t.Fatal(res)
}

func TestAny(t *testing.T) {
	data := []int{1, 2, 5}
	res := hadata.Any(data, func(x int) bool {
		return x < 4
	})
	if !res {
		t.Fatal("fucked")
	}
}

func TestCase(t *testing.T) {
	type TestStructXXX struct{}
	res := hadata.GetStructNameInFirstLetterLowerCase(TestStructXXX{})
	if res != "testStructXXX" {
		t.Fatal(res)
	}
}
