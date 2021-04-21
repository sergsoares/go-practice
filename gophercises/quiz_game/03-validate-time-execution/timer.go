package main

import (
	"os"
	"time"
	"fmt"
	"bufio"
)

const time_in_seconds = 3

func main() {
	timer := time.NewTimer(time.Second*time_in_seconds)

	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("\nTimer finished %d\n", time_in_seconds)
	}()

	fmt.Println("Started timer")
	reader :=	bufio.NewScanner(os.Stdin)

	fmt.Println("Press Key")
	reader.Scan()
	fmt.Println("KEY!!!")
}
