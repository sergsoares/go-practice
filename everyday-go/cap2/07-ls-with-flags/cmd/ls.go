package cmd

import (
	"fmt"
	"os"
)

func ExecuteLs(path string, withDir bool) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintln("Files in ", path)

	if withDir {
		output += "Name\tDirectory\t\n"
	} else {
		output += "Name\t\n"
	}

	for _, e := range files {
		if withDir {
			output += fmt.Sprintln(e.Name(), "\t", e.IsDir())
		} else {
			output += fmt.Sprintln(e.Name())
		}
	}

	return output, nil
}
