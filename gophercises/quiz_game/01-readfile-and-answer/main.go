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
	result    int
}

func main() {
	csvFile, err := os.Open("problems.csv")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("OK for Open file")

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewScanner(os.Stdin)

	for _, line := range csvLines {
		converted, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err)
		}

		item := problem{
			predicate: line[0],
			result:    converted,
		}

		fmt.Printf("Question %v: ", item.predicate)

		reader.Scan()
		value, _ := strconv.Atoi(reader.Text())

		if value != item.result {
			fmt.Println("Wrong Answer")
		} else {
			fmt.Println("Correct Answer")
		}
	}
}
