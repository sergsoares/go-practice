package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	var uri string

	flag.StringVar(&uri, "uri", "", "")
	flag.Parse()

	if uri == "" {
		fmt.Println("Not provided uri")
		os.Exit(1)
	}

	titleCh := make(chan string)
	errorCh := make(chan error)

	go titleOf(uri, titleCh, errorCh)

	select {
	case title := <-titleCh:
		fmt.Println("title", title)
	case err := <-errorCh:
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
}

func titleOf(uri string, titleCh chan string, errCh chan error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		errCh <- err
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errCh <- err
	}

	if res.Body != nil {
		defer res.Body.Close()

		res, _ := ioutil.ReadAll(res.Body)

		for _, l := range strings.Split(string(res), "\n") {
			if strings.Contains(l, "<title>") {
				titleCh <- l
				return
			}
		}
	}

	errCh <- fmt.Errorf("Title not found")
}
