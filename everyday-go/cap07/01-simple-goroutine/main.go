package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go func() {
		defer wg.Done()
		printLater("Hello", time.Millisecond*200)
	}()
	go func() {
		defer wg.Done()
		printLater("World", time.Millisecond*200)
	}()

	go func() {
		defer wg.Done()
		printLater(os.Getenv("USER"), time.Millisecond*200)
	}()

	wg.Wait()
}

func printLater(msg string, durantion time.Duration) {
	time.Sleep(durantion)
	fmt.Println(msg)
}
