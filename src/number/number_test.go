package number

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Added numbers positive", func(t *testing.T) {
		got := Add(3, 3)
		want := 6
		assertEqualNumber(t, got, want)
	})

	t.Run("Added numbers negative", func(t *testing.T) {
		got := Add(-20, -3)
		want := -23
		assertEqualNumber(t, got, want)
	})

	t.Run("Added numbers negative and positive", func(t *testing.T) {
		got := Add(20, -30)
		want := -10
		assertEqualNumber(t, got, want)
	})
}

func TestMulti(t *testing.T) {
	t.Run("Multi numbers positive", func(t *testing.T) {
		got := Multi(3, 5)
		want := 15
		assertEqualNumber(t, got, want)
	})

	t.Run("Multi numbers positive and negative", func(t *testing.T) {
		got := Multi(3, -5)
		want := -15
		assertEqualNumber(t, got, want)
	})

	t.Run("Multi numbers negative", func(t *testing.T) {
		got := Multi(-3, -3)
		want := 9
		assertEqualNumber(t, got, want)
	})

	t.Run("Multi number positive and zero", func(t *testing.T) {
		got := Multi(3, 0)
		want := 0
		assertEqualNumber(t, got, want)
	})
}

func assertEqualNumber(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
