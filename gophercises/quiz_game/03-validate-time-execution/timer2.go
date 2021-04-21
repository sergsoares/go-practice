//https://yourbasic.org/golang/time-reset-wait-stop-timeout-cancel-interval/

package main

import (
	"time"
	"fmt"
)

func main() {
	timeout := time.After(time.Second * 3)
	count := 1
	finish := make(chan bool)

	go func() {
		for {
			select {
				case <- timeout:
					fmt.Println("Timeout")
					finish <- true
					return
				default:
					fmt.Printf("Value %d", count)
					count++
					time.Sleep(time.Second*1)
			}
		}
	}()

	<- finish
	fmt.Println("Finish")
}

