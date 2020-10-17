package os

import "os"

func CreateIfNotExist(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
}
