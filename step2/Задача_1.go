package main

import (
	"os"
)

func ReadContent(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer f.Close()
	buffer := make([]byte, 1024)

	n, nerr := f.Read(buffer)
	if nerr != nil {
		return ""
	}

	return string(buffer[:n])

}
