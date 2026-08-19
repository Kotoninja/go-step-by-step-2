package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Data struct {
	body   float64
	status int
	err    error
}

func fetchMark(name string) *Data {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8082/mark?name=%s", name), nil)
	if err != nil {
		return &Data{body: 0, status: 500, err: err}
	}

	response, err := client.Do(req)
	if err != nil {
		return &Data{body: 0, status: response.StatusCode, err: err}
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return &Data{body: 0, status: response.StatusCode, err: err}
	}
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return &Data{body: 0, status: response.StatusCode, err: err}
	}
	return &Data{body: value, status: response.StatusCode, err: nil}
}

func Compare(name1, name2 string) (string, error) {
	mark1 := fetchMark(name1)
	if mark1.err != nil || mark1.status != 200 {
		return "", mark1.err
	}

	mark2 := fetchMark(name2)
	if mark2.err != nil || mark2.status != 200 {
		return "", mark2.err
	}

	if mark1.body > mark2.body {
		return ">", nil
	} else if mark1.body < mark2.body {
		return "<", nil
	} else {
		return "=", nil
	}

}
