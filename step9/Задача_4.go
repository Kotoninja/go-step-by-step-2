package main

import (
	"testing"
)

func TestAreAnagrams(t *testing.T) {
	s1, s2 := "кот", "ток"
	check := AreAnagrams(s1, s2)

	if !check {
		t.Error()
	}
	// different length

	check = AreAnagrams(s1+"1", s2)
	if check {
		t.Error()
	}
}
