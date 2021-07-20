package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	workItems := 5
	wg.Add(workItems)

	for i := 1; i <= workItems; i++ {
		go func(j int) {
			defer wg.Done()
			printLater(
				fmt.Sprint("Hello from ", j),
				time.Millisecond*100)
		}(i)
	}

	wg.Wait()

	wg.Add(workItems)
	for i := 1; i <= workItems; i++ {
		j := i
		go func() {
			defer wg.Done()
			printLater(
				fmt.Sprint("Hello from ", j),
				time.Millisecond*100)
		}()
	}

	wg.Wait()
}

func printLater(msg string, duration time.Duration) {
	time.Sleep(duration)
	fmt.Println(msg)
}
