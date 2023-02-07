package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Buffer implements Writer interface
	// We can use it to send it in as our Writer
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
