package main

import (
	"fmt"
	"log"
	"os"
	"tidy/redisb"
	"time"
)

func main() {

	f, _ := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()

	for {
		time.Sleep(time.Millisecond * 500)
		log.Println("Looping")

		allKeys, _ := redisb.GetAll()

		var result string
		for _, k := range allKeys {
			value, _ := redisb.Get(k)
			fmt.Println(k, value)
			result += fmt.Sprintln(k, value)

			redisb.Del(k)
		}
		f.WriteString(result)
	}
}
