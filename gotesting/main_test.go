package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expeceted result. Got: '%v', wanted: '%v'", got, expected)
	}
}

func TestSubstraction(t *testing.T) {
	got := 10 - 5
	expected := 5
	if got != expected {
		t.Errorf("Did not get expeceted result. Got: '%v', wanted: '%v'", got, expected)
	}
}
