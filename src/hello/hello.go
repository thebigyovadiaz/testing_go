package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	english = "English"

	englishHelloPrefix = "Hello"
	frenchHelloPrefix  = "Bonjour"
	spanishHelloPrefix = "Hola"
)

func Hello() string {
	return "Hello world!"
}

func Greeting(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s!", englishHelloPrefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = spanishHelloPrefix
	}

	return
}

func GreetingLanguage(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}
