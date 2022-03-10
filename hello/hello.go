package main

import "fmt"

type Language int

const (
	English Language = iota
	Spanish
)

func greetingPrefix(lang Language) (prefix string) {
	switch lang {
	case Spanish:
		prefix = "Hola"
	default:
		prefix = "Hello"
	}

	prefix = fmt.Sprintf("%s, ", prefix)
	return
}

func Hello(name string, lang Language) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(lang)

	return prefix + name
}

func main() {
	fmt.Println(Hello("", English))
}
