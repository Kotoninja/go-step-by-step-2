package main

import (
	"regexp"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
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

func filterByUser(tickets []Ticket, user string) []Ticket {
	result := []Ticket{}
	for _, t := range tickets {
		if t.User == user {
			result = append(result, t)
		}
	}
	return result
}

func filterByStatus(tickets []Ticket, status string) []Ticket {
	result := []Ticket{}
	for _, t := range tickets {
		if t.Status == status {
			result = append(result, t)
		}
	}
	return result
}

func GetTasks(text string, user *string, status *string) []Ticket {
	tickets := []Ticket{}

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		startIdx := strings.Index(line, "TICKET-")
		if startIdx == -1 {
			continue
		}

		ticketLine := line[startIdx:]

		array := getArray(ticketLine)

		if !checkDataText(array) {
			continue
		}

		timeStamp, _ := time.Parse("2006-01-02", strings.TrimSpace(array[3]))

		tickets = append(tickets, Ticket{
			Ticket: strings.TrimSpace(array[0]),
			User:   strings.TrimSpace(array[1]),
			Status: strings.TrimSpace(array[2]),
			Date:   timeStamp,
		})
	}

	if user != nil {
		tickets = filterByUser(tickets, *user)
	}

	if status != nil {
		tickets = filterByStatus(tickets, *status)
	}

	return tickets
}
