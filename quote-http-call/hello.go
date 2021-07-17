package main 

import (
	"rsc.io/quote"
	"fmt"
)

func Hello() string {
	return quote.Hello()
}

func main() {
	fmt.Printf(Hello())
	fmt.Printf("sergio")
}
