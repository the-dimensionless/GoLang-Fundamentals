package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	// Very custom logic for english bot
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// Very custom logic for english bot
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
