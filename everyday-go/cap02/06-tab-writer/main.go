package main

import (
	"fmt"
	"multiples-packages/cmd"
	"os"
	"text/tabwriter"
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

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, res)
	w.Flush()
}
