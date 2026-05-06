package main

import "fmt"
import "rsc.io/quote"

func main() {
	// read from stdin
	w := "Hello World"
	writeQuote(w)
}

func writeQuote(w string) {
	fmt.Println(quote.Hello())
}

func readInput() (string, error) {
	return "", nil
}
