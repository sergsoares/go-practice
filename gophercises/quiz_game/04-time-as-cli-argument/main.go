package main

import (
	"bufio"
	"time"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"flag"
)

type problem struct {
	predicate string
	result int
}

func main() {
	time_in_seconds := flag.Int("limit", 10, "Time in seconds")
  flag.Parse()
  fmt.Println("Timeout defined:", *time_in_seconds)

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

	timer := time.NewTimer(time.Second * time.Duration(*time_in_seconds))
	defer timer.Stop()

	go func(){
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
		<-timer.C
	}()

	<-timer.C
	fmt.Printf("You scored %d of %d", correct, total)
	defer csvFile.Close()
}