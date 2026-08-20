package main

import (
	"bufio"
	// "bytes"
	"context"
	"encoding/json"
	"errors"
	// "fmt"
	"io"
	"regexp"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string    `json:"ticket"`
	User   string    `json:"user"`
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

func checkDataText(data []string) bool {
	if len(data) != 4 {
		return false
	}

	pattern, err := regexp.MatchString(`^TICKET-\d+$`, strings.TrimSpace(data[0]))
	if !pattern || err != nil {
		return false
	}

	validStatuses := map[string]bool{
		"Готово":           true,
		"В работе":         true,
		"Не будет сделано": true,
	}
	if !validStatuses[strings.TrimSpace(data[2])] {
		return false
	}

	_, errTime := time.Parse("2006-01-02", strings.TrimSpace(data[3]))
	if errTime != nil {
		return false
	}

	return true
}

func getArray(text string) []string {
	return strings.Split(strings.TrimSpace(text), "_")
}

func getStudents(r io.Reader, user *string, status *string, chTicket chan []Ticket) {
	tickets := []Ticket{}

	newScaner := bufio.NewScanner(r)

	for newScaner.Scan() {
		line := strings.TrimSpace(newScaner.Text())

		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "TICKET-") {
			continue
		}

		array := getArray(line)

		if !checkDataText(array) {
			continue
		}

		timeStamp, _ := time.Parse("2006-01-02", strings.TrimSpace(array[3]))

		matchUser := user == nil || *user == array[1]
		matchStatus := status == nil || *status == array[2]

		if matchUser && matchStatus {
			tickets = append(tickets, Ticket{
				Ticket: array[0],
				User:   array[1],
				Status: array[2],
				Date:   timeStamp,
			})
		}
	}

	chTicket <- tickets
	return
}

func Writer(w io.Writer, chTicket chan []Ticket, chErr chan error) {
	tickets := <-chTicket

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(tickets); err != nil {
		chErr <- errors.New("write error")
		return
	}

	chErr <- nil
	return
}

func GetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	timeout time.Duration,
) error {
	chTicket := make(chan []Ticket, 1)
	// defer close(chTicket)

	chErr := make(chan error, 1)
	// defer close(chErr)

	// Timer and context
	ctxTimer, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	go getStudents(r, user, status, chTicket)
	go Writer(w, chTicket, chErr)

	select {
	case err := <-chErr:
		return err
	case <-ctxTimer.Done():
		return ctxTimer.Err()
	case <-time.After(timeout):
		return errors.New("Timeout") //fmt.Errorf("Timeout")
	}
}

// func main() {
// 	ctx := context.Background()

// 	input := strings.NewReader("TICKET-12345_Паша Попов_Готово_2024-01-01\nTICKET-12346_Иван Иванов_В работе_2024-01-02\nTICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03\nTICKET-12348_Паша Попов_В работе_2024-01-04")
// 	var output bytes.Buffer

// 	// username := nil
// 	status := "Готово"

// 	fmt.Println(GetTasks(ctx, input, &output, nil, &status, time.Minute))
// 	fmt.Println(output)
// }
