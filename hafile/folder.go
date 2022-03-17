package hafile

import (
	"io/ioutil"
	"os"

	"github.com/otiai10/copy"
)

//copy folder to another folder
func CopyFolder(src, dst string) error {
	err := copy.Copy(src, dst)
	return err
}

//check folder and create if not exist
func CheckFolder(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
		return true
	}
	return false
}

//check folder and create if not exist
func ExistFolder(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

//check file/folder in path folder
func CheckFileInPath(name string, path string) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, file := range files {
		if file.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

func RemFolder(path string) error {
	return os.RemoveAll(path)
}

func ListFolderFile(path string, remDir bool) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var ot []string
	for _, f := range files {
		if remDir && f.IsDir() {
			continue
		}
		ot = append(ot, f.Name())
	}
	return ot, nil
}
