package hafile

import (
	"encoding/xml"
	"io"
	"os"
)

type XmlFile struct {
	Path string
}

func NewXml(path string) *XmlFile {
	return &XmlFile{Path: path}
}

func (x *XmlFile) Parse(data interface{}) error {
	xmlFile, err := os.Open(x.Path)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, data)
	return err
}
