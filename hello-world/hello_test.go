package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Ingrid", "German")
		want := "Hallo, Ingrid"
		assertCorrectMessage(t, got, want)
	})
}

// Helper functions should accept `testing.TB` b/c
// it allows you to call helper functions from a test
// or a benchmark. It is an interface that both *testing.T
// and *testing.B(benchmark) satisfy
func assertCorrectMessage(t testing.TB, got, want string) {
	// This is needed so that when we fail the line is reported
	// from the test not the helper
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
