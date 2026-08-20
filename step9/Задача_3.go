package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "hello"
	if ReverseString(s) != "olleh" {
		t.Error()
	}
}
