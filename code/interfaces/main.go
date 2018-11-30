package main

import "fmt"

type bot interface {
	getGreeting() string //If any type with a function called 'getGreeting' and returns a string, it is now an honorary member of type bot
//It is a private function, all because it is lowercased! Other packages that imports this package can't access this function.
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
