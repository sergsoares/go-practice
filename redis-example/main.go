package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	var keys string
	flag.StringVar(&keys, "keys", "", "")
	flag.Parse()

	keysSplited := strings.Split(keys, ",")
	fmt.Println(keysSplited)
	saveKeys(keysSplited)

}

func saveKeys(keys []string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for k, v := range keys {
		fmt.Println(k, v)

		err := rdb.Set(ctx, fmt.Sprint("key", k), v, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}
