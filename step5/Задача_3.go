package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	URL        string
	Data       string
	StatusCode int   // код ответа
	Err        error // ошибка, если возникла
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	timeOut, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	client := &http.Client{}

	result := []*APIResponse{}

	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			result = append(result, &APIResponse{URL: url, Data: "", StatusCode: 500, Err: err})
			continue
		}
		response, err := client.Do(req.WithContext(timeOut))
		if err != nil {
			result = append(result, &APIResponse{URL: url, Data: "", StatusCode: 500, Err: context.DeadlineExceeded})
			continue
		}

		data, _ := io.ReadAll(response.Body)
		result = append(result, &APIResponse{URL: url, Data: string(data), StatusCode: response.StatusCode, Err: nil})
	}

	return result
}

// func main() {
// 	mCtx := context.Background()
// 	fmt.Println(FetchAPI(mCtx, []string{"https://ya.ru"}, 5*time.Microsecond))
// }
