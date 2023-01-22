package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Internal functions start w/ lowercase letters
// Prefix is initialized in the function as the "zero"
// value for the type (e.g. string "")
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// You don't need to specify the return value here
	// It is implied "prefix"
	return
}

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}
