package main

import (
	"fmt"

	"github.com/bitfield/script"
)

func main()  {
	fmt.Println("test")

	p := script.Exec("cd /est")

	if err := p.Error(); err != nil {
		//return fmt.Errorf("oh no: %w", err)
		fmt.Println(err)
	}

	fmt.Println(p.ExitStatus())
}