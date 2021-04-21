package main

import (
	"time"
	"fmt"
)

func main() {
	total_time := time.Duration(3) * time.Second
	finish := make(chan bool)

	index := 0
	be_executed := func() {
		fmt.Println(index)
		fmt.Println("finished")
		finish <- true
	}

	timer1 := time.AfterFunc(total_time, be_executed)
	defer timer1.Stop()

	fmt.Println("Star waiting")
  fmt.Println(index)

	<- finish
	index++
	index++
	index++
	index++

  fmt.Println(index)
}