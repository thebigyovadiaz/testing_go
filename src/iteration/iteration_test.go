package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat count", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertEqualMessage(t, got, want)
	})

	t.Run("repeat count", func(t *testing.T) {
		got := Repeat("ha", 8)
		want := "hahahahahahahaha"
		assertEqualMessage(t, got, want)
	})

	t.Run("repeat count negative", func(t *testing.T) {
		got := Repeat("pa", -8)
		want := "pa"
		assertEqualMessage(t, got, want)
	})
}

func TestIterArr(t *testing.T) {
	t.Run("iterating slice", func(t *testing.T) {
		arr := []int{3, 2, 4}
		got := IterArr(arr)
		want := 24
		assertEqualIntMessage(t, got, want)
	})

	t.Run("iterating empty slice", func(t *testing.T) {
		var arr []int
		got := IterArr(arr)
		want := 0
		assertEqualIntMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func assertEqualMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertEqualIntMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
