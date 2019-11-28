package trim

import (
	"bytes"
	"testing"
)

func TestMyTrim(t *testing.T) {
	input := []byte("    Hello    World!    ")
	want := []byte("Hello World!")
	got := myTrim(input)

	if bytes.Compare(got, want) {
		t.Errorf("Error: Got: %s, Want: %s", got, want)
	}
	// t.Fatal("not implemented")
}
