package main

import (
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
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

func BestStudents(names []string) (string, error) {
	var marks float64

	var students []Student
	for _, name := range names {
		data := fetchMark(name)
		if data.err != nil || data.status != 200 {
			return "", data.err
		}

		marks += data.body
		students = append(students, Student{name: name, mark: data.body})
	}

	students = slices.DeleteFunc(students, func(s Student) bool {
		return s.mark <= marks/float64(len(names))
	})
	names = []string{}
	for _, student := range students {
		names = append(names, student.name)
	}
	slices.Sort(names)
	return strings.Join(names, ","), nil
}
