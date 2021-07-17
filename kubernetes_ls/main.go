package main

import (
	"fmt"
	"flag"
)

func main() {
	useColor := flag.Bool("color", false, "display colors")
	types := flag.Bool("type", false, "display colors")

	flag.Parse()

	if *types == false {
		panic("NEED USE TYPE")
	}
	
	fmt.Println(*types)
	fmt.Println(*useColor)
}