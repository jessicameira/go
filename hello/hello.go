package main

import (
	"fmt"
	"rsc.io/quote"
	"example/greetings"
)

func main() {
	message:= greetings.Hello("User")
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	fmt.Println(message)
}