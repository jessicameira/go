package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"example/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Visiteur")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	fmt.Println(message)
}
