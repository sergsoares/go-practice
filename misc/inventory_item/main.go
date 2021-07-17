package main

import (
	"fmt"
)

type item struct {
	name string
	value string
}

func main() {
	fmt.Println("Initalized")

	a := item{
		name: "replicas",
		value: "34",
	}

	fmt.Println(a)
}