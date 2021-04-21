package main 

import (
	"fmt"
	"time"
)

const time_in_seconds = 3

func main() {
  timer := time.NewTimer(time.Second*time_in_seconds)

	defer timer.Stop()
	go func() {
		fmt.Printf("Done before")
		<-timer.C
	}()

	<-timer.C
	fmt.Printf("\nTimer finished %d\n", time_in_seconds)
}