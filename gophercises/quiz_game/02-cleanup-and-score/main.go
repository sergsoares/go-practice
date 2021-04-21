package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type problem struct {
	predicate string
	result int
}

func main() {
	csvFile, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println(err)
	}


	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewScanner(os.Stdin)

	correct := 0
	total := 0
	for index, line := range csvLines {
		index++
		total++

		converted, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err)
		}

		item := problem{
			predicate: line[0],
			result: converted,
		}

		fmt.Printf("Question #%d (%v): ", index, item.predicate)

		reader.Scan()
		value, _:= strconv.Atoi(reader.Text())

		if value == item.result {
			correct++
		}
	}

	fmt.Printf("You scored %d of %d", correct, total)
	defer csvFile.Close()
}