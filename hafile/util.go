package hafile

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/johnhaha/hakit/hadata"
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

// check if file exist
func ExistFile(path string) (exist bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// check if file exist, create if not exist
func CheckFile(path string) (exist bool) {
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

// write to specific line
func WriteLine(path string, line int, content string) error {
	if line < 0 {
		return errors.New("no such line")
	}
	check := ExistFile(path)
	if !check {
		return errors.New("no such file")
	}
	input, err := os.ReadFile(path)
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
	err = os.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}

// append
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

func ReplaceFileInDir(path string, from string, to string, neglect ...string) error {
	err := filepath.Walk(path, func(path string, info fs.FileInfo, _ error) error {
		if !info.IsDir() && !hadata.IsInSlice(neglect, info.Name()) {
			writer := NewFileWriter(path)
			err := writer.Replace(from, to)
			if err != nil {
				log.Print(err)
			}
		}
		return nil
	})
	return err
}

// will rem file or folder if name contains 'name'
func RemFileInDirByName(path string, name string, onRem func(string)) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, f := range files {
		if n := f.Name(); strings.Contains(n, name) {
			err = os.RemoveAll(path + "/" + n)
			if err != nil {
				return err
			}
			onRem(n)
		}
	}
	return nil
}
