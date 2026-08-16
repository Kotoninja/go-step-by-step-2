package main

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

