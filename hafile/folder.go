package hafile

import "github.com/otiai10/copy"

func CopyFolder(src, dst string) error {
	err := copy.Copy(src, dst)
	return err
}
