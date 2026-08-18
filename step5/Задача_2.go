// package main

// import (
// 	"context"
// 	"io"
// 	"net/http"
// 	"time"
// )

// type APIResponse struct {
// 	Data       string
// 	StatusCode int
// }

// func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
// 	timeOutCtx, cancel := context.WithTimeout(ctx, timeout)
// 	defer cancel()

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response, err := client.Do(req.WithContext(timeOutCtx))
// 	if err != nil {
// 		return nil, context.DeadlineExceeded
// 	}
// 	defer response.Body.Close()

// 	data, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &APIResponse{Data: string(data), StatusCode: response.StatusCode}, nil
// }
