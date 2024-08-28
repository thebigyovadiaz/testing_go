package main

import (
	"fmt"
	"github.com/thebigyovadiaz/testing_go/src/mocking"
	"os"
	"time"
)

func main() {
	fmt.Println("Learning TDD with Go")
	// _ = http.ListenAndServe(":5001", http.HandlerFunc(dependency_injection.MyGreetHandler))

	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, Slept: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
