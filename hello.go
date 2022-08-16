package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", spanish))
}
