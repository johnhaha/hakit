package hakit_test

import (
	"testing"

	"github.com/johnhaha/hakit/hacmd"
	"github.com/johnhaha/hakit/hafile"
)

func TestCopy(t *testing.T) {
	err := hafile.Copy("./README.md", "hafile/README.md")
	if err != nil {
		t.Fatal(err)
	}
}

//test check and create file
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
