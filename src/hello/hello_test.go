package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello world!"
	validCorrectMessage(t, got, want)
}

func TestGreeting(t *testing.T) {
	t.Run("Saying greeting to people", func(t *testing.T) {
		got := Greeting("Delia")
		want := "Hello, Delia!"
		validCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello world!' when name is an empty string", func(t *testing.T) {
		got := Greeting("")
		want := "Hello, World!"
		validCorrectMessage(t, got, want)
	})
}

func TestGreetingLanguage(t *testing.T) {
	t.Run("Greeting in Spanish", func(t *testing.T) {
		got := GreetingLanguage("Delia", "Spanish")
		want := "Hola, Delia!"
		validCorrectMessage(t, got, want)
	})

	t.Run("Greeting in French", func(t *testing.T) {
		got := GreetingLanguage("Delia", "French")
		want := "Bonjour, Delia!"
		validCorrectMessage(t, got, want)
	})

	t.Run("Greeting in English", func(t *testing.T) {
		got := GreetingLanguage("Delia", "English")
		want := "Hello, Delia!"
		validCorrectMessage(t, got, want)
	})

	t.Run("Greeting without name in English", func(t *testing.T) {
		got := GreetingLanguage("", "English")
		want := "Hello, World!"
		validCorrectMessage(t, got, want)
	})
}

func validCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
