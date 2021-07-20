package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func main() {
	x := 5
	y := 5
	fmt.Printf("%d + %d = %d\n", x, y, sum(x, y))
}
