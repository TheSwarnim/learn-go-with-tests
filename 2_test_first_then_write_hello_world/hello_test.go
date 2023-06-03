package main

import "testing"

func TestHello(t *testing.T) {
	// rewrite of 1_st iteration
	// got := Hello("Swarnim")
	// want := "Hello, Swarnim!"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	// want = "Hello, World!"

	// if got == want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Swarnim", "")
		want := "Hello, Swarnim!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	// It is important that your tests are clear specifications of what the code needs to do.
	// But there is repeated code when we check if the message is what we expect.

	// new test for Spanish
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Swarnim", "Spanish")
		want := "Hola, Swarnim!"
		assertCorrectMessage(t, got, want)
	})

	// new test for French
	t.Run("in French", func(t *testing.T) {
		got := Hello("Swarnim", "French")
		want := "Bonjour, Swarnim!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	// This will help other developers track down problems easier.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}