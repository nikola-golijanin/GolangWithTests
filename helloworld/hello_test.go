package helloworld

import "testing"

func TestHello(t *testing.T) {

	t.Run("Say hello to Nikola", func(t *testing.T) {
		got := Hello("Nikola", "")
		want := "Hello, Nikola"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello world when name is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
