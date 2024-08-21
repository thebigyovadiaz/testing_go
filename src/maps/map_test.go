package maps

import "testing"

var dictionary = Dictionary{"test": "this is just a test"}

func TestSearch(t *testing.T) {
	t.Run("key test is exist", func(t *testing.T) {
		key := "test"
		got := dictionary.Search(key)
		want := "this is just a test"
		assertString(t, got, want, key)
	})

	t.Run("key test2 does not exist", func(t *testing.T) {
		key := "test2"
		got := dictionary.Search(key)
		want := "key does not exist"
		assertString(t, got, want, key)
	})
}

func assertString(t testing.TB, got, want, key string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, key)
	}
}
