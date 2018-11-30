package main

import "fmt"

type englishBot_2 struct {}

func (englishBot_2) getGreeting() string {
	return "Hello"
}

func printGreeting_2(eb englishBot_2) {
	fmt.Println(eb.getGreeting())
}

type spanishBot_2 struct {}

func (spanishBot_2) getGreeting() string {
	return "Halo"
}

func printGreeting_3(sb spanishBot_2) { //go does not even allow func name overloading
	fmt.Println(sb.getGreeting())
}

func main_2() {
	eb := englishBot_2{}
	sb := spanishBot_2{}
	printGreeting_2(eb)
	printGreeting_3(sb)
}