package hafile

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type Filer struct {
	FileWriter
	FileReader
}

func NewFiler(path string) *Filer {
	CheckFile(path)
	return &Filer{
		FileWriter: FileWriter{Path: path},
		FileReader: FileReader{Path: path},
	}
}

type FileWriter struct {
	Path string
}

type FileReader struct {
	Path string
}

func NewFileWriter(path string) *FileWriter {
	CheckFile(path)
	return &FileWriter{Path: path}
}

func NewFileReader(path string) *FileReader {
	CheckFile(path)
	return &FileReader{Path: path}
}

//find text in line
func (reader *FileReader) FineText(text string) (int, string, error) {
	f, err := os.Open(reader.Path)
	if err != nil {
		return -1, "", err
	}
	defer f.Close()
	line := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			return line, scanner.Text(), nil
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		return -1, "", err
	}
	return -1, "", nil
}

//read specific line in file
func (reader *FileReader) ReadLine(line int) (string, error) {
	fileIO, err := os.OpenFile(reader.Path, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	sc := bufio.NewScanner(fileIO)
	l := 1
	for sc.Scan() {
		if l == line {
			return sc.Text(), sc.Err()
		}
		l++
	}
	return "", errors.New("not found")
}

//read specific line in file
func (reader *FileReader) ReplaceLine(line int, content string) error {
	err := WriteLine(reader.Path, line, content)
	return err
}

//read file in string
func (reader *FileReader) Read() (string, error) {
	data, err := ioutil.ReadFile(reader.Path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (writer *FileWriter) Write(text string) error {
	d := []byte(text)
	err := os.WriteFile(writer.Path, d, 0644)
	return err
}

func (writer *FileWriter) Update(text string) error {
	f, err := os.OpenFile(writer.Path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString("\n" + text + "\n"); err != nil {
		return err
	}
	return nil
}

func (writer *FileWriter) Replace(from string, to string) error {
	input, err := ioutil.ReadFile(writer.Path)
	if err != nil {
		return err
	}

	output := bytes.Replace(input, []byte(from), []byte(to), -1)

	if err = ioutil.WriteFile(writer.Path, output, 0644); err != nil {
		return err
	}
	return nil
}
