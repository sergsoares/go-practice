package main

import (
	"os"
	"fmt"
	"text/template"
)

type Welcome struct {
	Name string
	Message string
}

func main() {
	fmt.Println("Template Generated")

	td := Welcome{"Sergio", "keep growing!!!"}
	t, err := template.New("welc").Parse("Welcome {{ .Name}}, your message is {{ .Message}}")

	if err != nil {
		panic("Error for templates")
	}

	err = t.Execute(os.Stdout, td)

	if err != nil {
		panic("Error for parse")
	}
}