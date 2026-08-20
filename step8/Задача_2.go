package main

import (
	"context"
	"io"
	"os"
	"time"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	ctxTime, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	defer close(result)
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	dataChan := make(chan []byte, 1)

	go func() {
		data, err := io.ReadAll(file)
		if err != nil {
			return
		}
		dataChan <- data
	}()

	select {
	case res := <-dataChan:
		result <- res
	case <-ctxTime.Done():
		return
	}
}
