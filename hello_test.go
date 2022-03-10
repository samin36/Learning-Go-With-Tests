package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", English)
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Angela", Spanish)
		want := "Hola, Angela"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string in Spanish", func(t *testing.T) {
		got := Hello("", Spanish)
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
}
