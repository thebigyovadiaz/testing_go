package dependency_injection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s\n", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World!")
}
