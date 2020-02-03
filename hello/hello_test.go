package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("kunle", "")
		want := "Hello, kunle"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Rafael", "Spanish")
		want := "Hola, Rafael"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, World' when an empty string is supplied in Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Mendy", "French")
		want := "Bonjour, Mendy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Bonjour, World' when an empty string is supplied in French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"

		assertCorrectMessage(t, got, want)
	})
}
