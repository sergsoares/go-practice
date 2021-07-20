package cmd

import (
	"fmt"
	"os"
)

func ExecuteLs(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintln("Files in ", path)
	output += "Name\tDirectory\t\n"

	for _, e := range files {
		output += fmt.Sprintln(e.Name(), "\t", e.IsDir())
	}

	return output, nil
}
