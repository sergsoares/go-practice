package main

import (
	"flag"
	"fmt"
)

type values struct {
	x int
	y int
}

func sum(v values) int {
	return v.x + v.y
}

func parseArgumentsToValues() values {
	var x int
	var y int
	flag.IntVar(&x, "x", 0, "teste")
	flag.IntVar(&y, "y", 0, "teste")
	flag.Parse()

	v := values{
		x: x,
		y: y,
	}
	return v
}

func main() {
	v := parseArgumentsToValues()

	fmt.Printf("%d + %d = %d\n", v.x, v.y, sum(v))
}
