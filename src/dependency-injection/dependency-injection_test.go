package dependency_injection

import (
	"fmt"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("greet string with injection", func(t *testing.T) {
		stdout := os.Stdout
		Greet(stdout, "Chris")

		want := "Hello, Chris"
		fmt.Println(want)

		/*if got != want {
			t.Errorf("got %q, want %q", got, want)
		}*/
	})
}
