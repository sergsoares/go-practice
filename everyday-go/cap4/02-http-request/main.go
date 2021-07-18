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
	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	log.Println("Preparing request")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, err := spaceClient.Do(req)
	log.Println("Request Done")
	log.Println(res.Status)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("body already readed")

	p := people{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Number)

}
