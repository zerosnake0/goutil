package os

import (
	"fmt"
	"os"
)

func FileExists(filename string) (bool, error) {
	return exists(filename, false)
}

func FolderExists(folder string) (bool, error) {
	return exists(folder, true)
}

func exists(path string, dir bool) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() != dir {
			if dir {
				return false, fmt.Errorf("%q is not a directory", path)
			}
			return false, fmt.Errorf("%q is not a file", path)
		}
		return true, nil
	}
	if !os.IsNotExist(err) {
		return false, err
	}
	return false, nil
}
