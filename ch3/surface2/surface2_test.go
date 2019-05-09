package main

import "testing"

func TestMyRound(t *testing.T) {
	input := 99.99999
	want := 99.999
	got := myRound(input, 3)
	if got != want {
		t.Errorf("Expected %f, but got %f", got, want)
	}
}

func TestMyRoundZero(t *testing.T) {
	input := 99.99999
	want := 99.0
	got := myRound(input, 0)
	if got != want {
		t.Errorf("Expected %f, but got %f", got, want)
	}
}
