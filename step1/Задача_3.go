package main

import (
	"errors"
	"strings"
)

type UpperWriter struct {
	UpperString string
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func (obj *UpperWriter) Write(p []byte) (n int, err error) {
	obj.UpperString = strings.ToTitle(string(p))
	return len(p), nil
}
