package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	result := Contains(numbers, 4)

	if !result {
		t.Error()
	}

	result = Contains(numbers, 5)

	if result {
		t.Error()
	}
}
