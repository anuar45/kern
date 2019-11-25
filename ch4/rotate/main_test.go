package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5}
	want := []int{2, 3, 4, 5, 0, 1}
	got := rotate(input, 2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error. Want: %v, Got: %v", want, got)
	}

}
