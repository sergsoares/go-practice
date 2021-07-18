package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USER")
	fmt.Println("Username is:", name)
}
