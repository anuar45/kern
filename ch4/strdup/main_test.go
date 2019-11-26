package main

import (
	"reflect"
	"testing"
)

func TestStrDup(t *testing.T) {
	input := []string{
		"january",
		"january",
		"january",
		"january",
		"january",
		"february",
		"february",
		"february",
		"march",
		"march",
		"april",
		"april",
		"april",
	}

	want := []string{
		"january",
		"february",
		"march",
		"april",
	}

	got := strDup(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Error: Want: %v, Got: %v", want, got)
	}
}
