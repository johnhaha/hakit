package hafile

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func OpenFile(path string) *os.File {
	dir := filepath.Dir(path)
	CheckFolder(dir)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return nil
	}
	return file
}

//check if file exist
func ExistFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

//check if file exist
func CheckFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		return false
	}
	return true
}

//write to specific line
func WriteLine(path string, line int, content string) error {
	if line < 0 {
		return errors.New("no such line")
	}
	check := ExistFile(path)
	if !check {
		return errors.New("no such file")
	}
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	if len(lines) < line-1 {
		for i := len(lines); i < line; i++ {
			lines = append(lines, "")
		}
	}
	lines[line-1] = content
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}

//append
func AppendLine(path string, content string) error {
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		return err
	}
	return nil
}