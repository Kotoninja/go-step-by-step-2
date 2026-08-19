package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://localhost:8081/provideData", nil)
		if err != nil {
			fmt.Fprint(w, err)
		}

		response, err := client.Do(req)
		if err != nil {
			fmt.Fprint(w, context.DeadlineExceeded)
		}

		data, _ := io.ReadAll(response.Body)
		w.Write(data)
	})

	timeoutHandler := http.TimeoutHandler(handler, maxTimeout, "")

	http.Handle("/readSource", timeoutHandler)
	http.ListenAndServe(":8080", nil)
}

// func Eight() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello world!")
// 	})
// 	http.ListenAndServe(":8081", nil)
// }

// func main() {
// 	go Eight()
// 	select {}
// }
