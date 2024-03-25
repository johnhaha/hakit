package hakit_test

import (
	"testing"

	"github.com/johnhaha/hakit/hacmd"
	"github.com/johnhaha/hakit/haerr"
	"github.com/johnhaha/hakit/hafile"
	"github.com/johnhaha/hakit/hareq"
)

func TestCopy(t *testing.T) {
	err := hafile.Copy("./README.md", "hafile/README.md")
	if err != nil {
		t.Fatal(err)
	}
}

// test check and create file
func TestCreateFile(t *testing.T) {
	hafile.CheckFile("test.yml")
}

func TestWriteToFile(t *testing.T) {
	hafile.WriteLine("test", 10, "changed")
}

func TestRelative(t *testing.T) {
	type Case struct {
		Input  string
		Output bool
	}
	cases := []Case{
		{Input: "path", Output: true},
		{Input: "/path", Output: false},
	}
	for _, c := range cases {
		res := hafile.CheckRelativePath(c.Input)
		if res != c.Output {
			t.Fatal("failed")
		}
	}
}

func TestRun(t *testing.T) {
	ot, err := hacmd.Run("pwd")
	if err != nil {
		t.Fatal(err)
	}
	if ot != "/Users/junwu/develop/pkg/hakit" {
		t.Fatal("🤬" + ot + "nnnn")
	}
}

func TestUpdateFile(t *testing.T) {
	wt := hafile.NewFileWriter("test")
	wt.Update("okddddokok")
}

func TestFindText(t *testing.T) {
	fileReader := hafile.NewFileReader("test")
	x, _, err := fileReader.FineText("okddddokok")
	if err != nil {
		t.Fatal(err)
	}
	if x != 11 {
		t.Fatal(x)
	}
}

func TestReadLine(t *testing.T) {
	fileReader := hafile.NewFileReader("test")
	res, err := fileReader.ReadLine(11)
	if err != nil {
		t.Fatal(err)
	}
	if res != "okddddokok" {
		t.Fatal(res)
	}
}

func BenchmarkReadLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileReader := hafile.NewFileReader("test")
		fileReader.ReadLine(11)

	}
}

func TestDownloadFile(t *testing.T) {
	err := hareq.DownloadFileFromUrl("https://www.codegrepper.com/codeimages/user-input-golang.png", ".", "xxx.png")
	if err != nil {
		t.Fatal("failed")
	}
}

func TestCheckInPath(t *testing.T) {
	check, err := hafile.CheckFileInPath("hafile", "/Users/junwu/develop/pkg/hakit")
	if err != nil {
		t.Fatal(err)
	}
	if !check {
		t.Fatal("failed")
	}
}

func TestReplaceText(t *testing.T) {
	writer := hafile.NewFileWriter("/Users/junwu/develop/pkg/hakit/test")
	err := writer.Replace("replace", "dog")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReplaceInDir(t *testing.T) {
	err := hafile.ReplaceFileInDir("/Users/junwu/develop/proj/chess", "dog", "cat", "test2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRem(t *testing.T) {
	hafile.RemFolder("/Users/junwu/develop/pkg/hakit/testssss")
}

func TestListFolderFile(t *testing.T) {
	res, err := hafile.ListFolderFile("hafile", false)
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 6 {
		t.Fatal(res)
	}
}

func TestRemFile(t *testing.T) {
	err := hafile.RemFileInDirByName("test", "1", func(s string) {})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTrace(t *testing.T) {
	haerr.Trace()
	t.Fatal("res")
}

func TestParseCsv(t *testing.T) {
	type Data struct {
		Id   string `csv:"client_id"`
		Name string `csv:"client_name"`
		Age  string `csv:"client_age"`
	}
	data := []*Data{}
	err := hafile.ParseCsv("/Users/junwu/develop/pkg/hakit/test/test.csv", &data)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(data[0].Name)
}
