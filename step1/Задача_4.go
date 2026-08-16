package main

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	copyLen := int(n)
	if copyLen == 0 || copyLen > len(data) {
		copyLen = len(data)
	}

	_, errWrite := w.Write(data[:copyLen])

	if errWrite != nil {
		return errWrite
	}

	return nil
}
