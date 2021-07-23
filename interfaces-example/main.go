package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %s article was written by %s", a.Title, a.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Great test about Go",
		Author: "Sammy dark",
	}
	Print(a)
}
func Print(a Stringer) {
	fmt.Println(a.String())
}
