/*
Target: List most logged IP from logs file.
*/

package main

import(
	"log"
	"os"
	"bufio"
	"fmt"
	"strings"
)

var path = "../generated.log"

func main() {
	data, err := os.Open(path)

	if err != nil {
		log.Fatal("Error Reading file", err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	var ips map[string]int
	ips = make(map[string]int)

	var count int

	for scanner.Scan() {
		ip := strings.Split(scanner.Text(), " ")[0]

		if _, ok := ips[ip]; ok {
			ips[ip]++
		} else {
			ips[ip] = 1
		}

		count++
	}

	var top_position string
	top_position = ""

	for k, v := range ips {
		if top_position == "" {
			top_position = k
		}

		if v > ips[top_position] {
			top_position = k
		}
	}

	fmt.Println("Top Value:", top_position, "with", ips[top_position], "records." )
}