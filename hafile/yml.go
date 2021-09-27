package hafile

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type YmlReader struct {
	Path string
}

func NewYmlReader(path string) *YmlReader {
	return &YmlReader{Path: path}
}

type YmlWriter struct {
	Path string
}

func NewYmlWriter(path string) *YmlWriter {
	return &YmlWriter{Path: path}
}

func (reader *YmlReader) Parser(data interface{}) error {
	yfile, err := ioutil.ReadFile(reader.Path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yfile, data)
	if err != nil {
		return err
	}
	return nil
}

func (writer *YmlWriter) Write(data interface{}) error {
	datao, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(writer.Path, datao, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (writer *YmlWriter) Append(data interface{}) error {
	datao, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(writer.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write(datao)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
