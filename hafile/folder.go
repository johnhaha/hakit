package hafile

import (
	"os"

	"github.com/otiai10/copy"
)

//copy folder to another folder
func CopyFolder(src, dst string) error {
	err := copy.Copy(src, dst)
	return err
}

func CheckFolder(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
		return true
	}
	return false
}
