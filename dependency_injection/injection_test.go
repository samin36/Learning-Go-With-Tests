package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "World")

	got := buffer.String()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
