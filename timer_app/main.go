
// Activity to kind with slices a little and string manipulation.

package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	var time_now time.Time = time.Now()

	main_time_blocks := strings.Split(time_now.String(), " ")
	date_blocks := strings.Split(main_time_blocks[0], "-")
	hour_block := strings.Split(main_time_blocks[1], ":")

	fmt.Println("Year:", date_blocks[0])
	fmt.Println("Month:", date_blocks[1])
	fmt.Println("Day:", date_blocks[2])
	fmt.Println("Hour:", hour_block[0])
	fmt.Println("Minute:", hour_block[1])
	fmt.Println("Second:", hour_block[2])
	fmt.Println("Miliseconds:", hour_block[2])
}

//https://yourbasic.org/golang/split-string-into-slice/