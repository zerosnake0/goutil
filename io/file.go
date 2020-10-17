package io

import (
	"io"
	"os"
)

func AppendFile(filename string, cb func(w io.Writer) error) error {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	pos, err := fp.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	if pos > 0 {
		_, err := fp.Seek(-1, io.SeekEnd)
		if err != nil {
			return err
		}
		var b [1]byte
		_, err = fp.Read(b[:])
		if err != nil {
			return err
		}
		if b[0] != '\n' {
			fp.Write([]byte{'\n'})
		}
	}
	return cb(fp)
}
