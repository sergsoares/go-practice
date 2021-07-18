package main

import (
	"fmt"
	"multiples-packages/cmd"
	"os"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to get work directoruy: %s", err.Error())
		panic("")
	}

	res, err := cmd.ExecuteLs(wd)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to list files %s, error: %s", wd, err.Error())
		panic("")
	}

	fmt.Println(res)
}
