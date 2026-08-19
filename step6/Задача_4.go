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

type Student struct {
	name string
	mark float64
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

func CompareList(names []string) (map[string]string, error) {
	studentsMap := map[string]float64{}

	var marks float64

	for _, name := range names {
		mark := fetchMark(name)
		if mark.err != nil || mark.status != 200 {
			return map[string]string{}, mark.err
		}

		studentsMap[name] = mark.body
		marks += mark.body
	}

	studentsMapWithSymbol := map[string]string{}
	averageMark := float64(marks / float64(len(names)))

	for student, mark := range studentsMap {
		if mark > averageMark {
			studentsMapWithSymbol[student] = ">"
		} else if mark == averageMark {
			studentsMapWithSymbol[student] = "="
		} else {
			studentsMapWithSymbol[student] = "<"
		}
	}
	return studentsMapWithSymbol, nil
}
