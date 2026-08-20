package main

import (
	"strings"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	var result int

	for i := 0; i < len(questions); i++ {
		select {
		case res := <-answerCh:
			if strings.EqualFold(strings.ToLower(strings.TrimSpace(res)), strings.ToLower(strings.TrimSpace(answers[i]))) {
				result++
			}
		case <-time.After(time.Second):
			continue
		}
	}
	return result
}
