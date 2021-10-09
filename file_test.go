package hakit_test

import (
	"testing"

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
