package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := errgroup.Group{}

	g.Go(func() error {
		printLater("Sergio", time.Millisecond*800)
		return nil
	})

	g.Go(func() error {
		printLater("test", time.Millisecond*300)
		return fmt.Errorf("everFailed\n")
	})

	g.Go(func() error {
		// printLater(os.Getenv("USER"), time.Millisecond*100)
		return fmt.Errorf("get user failed\n")
		// return nil
	})

	err := g.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error found: %s", err.Error())
		os.Exit(1)
	}

	fmt.Print("All Work completed\n")
}

func printLater(msg string, duration time.Duration) {
	time.Sleep(duration)
	fmt.Println(msg)
}
