package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type Italian struct{}

type Portuguese struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(n string) string {
	return fmt.Sprintf("Ciao %s!", n)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(n string) string {
	return fmt.Sprintf("Olá %s!", n)
}

func SayHello(visitorName string, anything Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", anything.LanguageName(), anything.Greet(visitorName))
}
