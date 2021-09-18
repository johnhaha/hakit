package hafile

import "github.com/otiai10/copy"

//copy folder to another folder
func CopyFolder(src, dst string) error {
	err := copy.Copy(src, dst)
	return err
}
