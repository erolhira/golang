package main

import "fmt"

type composedInterface interface {
	//like in Java:
	//interface composedInterface extends readable, writeable
	readable
	writeable
}
type readable interface {
	read() string
}
type writeable interface {
	write(string)
}

type bot interface {
	getGreeting() string
}
type englishBot struct {}
type spanishBot struct {}

func testInterfaces() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() (string) {
	return "Hi there"
}

func (spanishBot) getGreeting() (string) {
	return "Hola"
}