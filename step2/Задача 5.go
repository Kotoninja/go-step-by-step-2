package main

import (
	"bufio"
	// "fmt"
	"os"
	"strings"
	"time"
	"errors"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	f, _ := os.Open(inputFileName)

	var logs []string

	buffer := bufio.NewScanner(f)

	for buffer.Scan() {
		line := strings.Split(strings.TrimSpace(buffer.Text()), " ")
		if len(line) > 0 {
			date, err := time.Parse("02.01.2006", line[0])
			if err != nil {
			}
			if (date.After(start) || date.Equal(start)) && (date.Before(end) || date.Equal(end)) {
				logs = append(logs, buffer.Text())
			}
		}
	}
	if len(logs) != 0 {
    	return logs, nil
    }
    return logs, errors.New("Not logs")
}

// func main() {
// 	fmt.Println(ExtractLog("/home/user/code/go-step-by-step-2/step2/text.txt", time.Date(2022, 12, 14, 0, 0, 0, 0, time.UTC), time.Date(2022, 12, 15, 0, 0, 0, 0, time.UTC)))
// }
