package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type ScrapeRun struct {
	Results map[string]ScrapeResult
	Lock    *sync.Mutex
}

type ScrapeResult struct {
	HTTPCode   int
	HTTPLength int64
	Error      error
}

func scrape(url string, sr *ScrapeRun) {
	res, err := http.Get(url)
	if err != nil {
		sr.Lock.Lock()
		defer sr.Lock.Unlock()

		sr.Results[url] = ScrapeResult{
			Error:    err,
			HTTPCode: http.StatusBadGateway,
		}
		return
	}

	defer res.Body.Close()

	sr.Lock.Lock()
	defer sr.Lock.Unlock()

	length := res.ContentLength
	if length == -1 {
		body, _ := ioutil.ReadAll(res.Body)
		length = int64(len(body))
	}

	sr.Results[url] = ScrapeResult{
		HTTPCode:   res.StatusCode,
		HTTPLength: length,
	}
}

func main() {
	url1 := "https://inlets.dev"
	url2 := "https://openfaas.com"
	url3 := "https://openfas.com"

	sr := ScrapeRun{
		Lock:    &sync.Mutex{},
		Results: make(map[string]ScrapeResult),
	}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		scrape(url1, &sr)
	}()

	go func(u string) {
		defer wg.Done()
		scrape(u, &sr)
	}(url2)

	go func(u string) {
		defer wg.Done()
		scrape(u, &sr)
	}(url3)

	wg.Wait()

	for k, v := range sr.Results {
		if v.Error != nil {
			fmt.Printf("- %s = error: %s\n", k, v.Error.Error())
		} else {
			fmt.Printf("- %s = [%d] %d bytes\n", k, v.HTTPCode, v.HTTPLength)
		}
	}
}
