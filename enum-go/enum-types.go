package main

import "fmt"

type Season int64

const (
	Undefined Season = iota
	Summer
	Autumn
	Winter
	Spring
)

func (s Season) String() string {
	switch s {
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	case Spring:
		return "spring"
	default:
		return "undefined"
	}
}

func main() {
	fmt.Println(Summer)
	fmt.Println(Spring)

	printSeason(Autumn)

	var se Season = 132
	printSeason(se)
	if se == Undefined {
		fmt.Println("logged")
	} else {
		fmt.Println("not undefined")
	}

}

func printSeason(s Season) {
	fmt.Println("season:", s)
}
