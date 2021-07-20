package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:"number"`
}

func main() {
	apiURL := "http://api.open-notify.org/astros.json"

	p, err := getAstronauts(apiURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of people in space", p.Number)
}

func getAstronauts(apiUrl string) (people, error) {
	var p people
	httpClient := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)

	if err != nil {
		return p, err
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")
	res, err := httpClient.Do(req)
	if err != nil {
		return p, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p, err
	}

	if err := json.Unmarshal([]byte(body), &p); err != nil {
		return p, err
	}

	return p, nil
}
