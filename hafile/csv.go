package hafile

import (
	"os"

	"github.com/gocarina/gocsv"
)

func ParseCsv(path string, data any) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := gocsv.UnmarshalFile(file, data); err != nil {
		return err
	}
	return nil
}
