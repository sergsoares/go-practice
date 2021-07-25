package main

import (
	"flag"
	"fmt"
	"strings"
	"tidy/redisb"
)

func main() {
	var keys string
	flag.StringVar(&keys, "keys", "", "")
	flag.Parse()

	keysSplited := strings.Split(keys, ",")
	fmt.Println(keysSplited)

	for k, v := range keysSplited {
		fmt.Println(k, v)
		err := redisb.Set(fmt.Sprint("key", k), v)
		if err != nil {
			panic(err)
		}
	}
}
