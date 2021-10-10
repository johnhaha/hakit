package hafile

import (
	"errors"
	"os"
)

//get current relative path's full path
func FullPath(path string) (string, error) {
	check := CheckRelativePath(path)
	if !check {
		return "", errors.New("path not correct")
	}
	fullPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fullPath + "/" + path, nil
}

func CheckRelativePath(path string) bool {
	if path == "" {
		return false
	}
	start := path[:1]
	return start != "/"
}
