package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	lock := sync.RWMutex{}
	data := make(map[string]string)
	data["key"] = "value"

	wg.Add(2)

	go readOperation(data, "key", &lock)
	go writeOperation(data, "key", "bla", &lock)

	wg.Wait()

}

func readOperation(data map[string]string, key string, lock *sync.RWMutex) {
	defer wg.Done()
	lock.RLock()
	defer lock.RUnlock()
	fmt.Println(data["key"])
}

func writeOperation(data map[string]string, key, value string, lock *sync.RWMutex) {
	time.Sleep(time.Millisecond * 400)
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	data[key] = value
}
