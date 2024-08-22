package main

import (
	"fmt"
	dependency_injection "github.com/thebigyovadiaz/testing_go/src/dependency-injection"
	"net/http"
)

func main() {
	fmt.Println("Learning TDD with Go")
	_ = http.ListenAndServe(":5001", http.HandlerFunc(dependency_injection.MyGreetHandler))
}
