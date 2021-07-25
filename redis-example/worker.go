package main

import (
	"fmt"
	"os"
	"tidy/redisb"
	"time"
)

func main() {

	os.Create("result.txt")

	var result string
	for {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Timer")

		for k, v := range redisb.GetAll() {
			fmt.Println(k, v)
			result += fmt.Sprint(k, v)
		}
		os.WriteFile("result.txt", []byte(result), os.ModeAppend)
	}

}
