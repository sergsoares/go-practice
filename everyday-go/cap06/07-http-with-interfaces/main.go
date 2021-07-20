package main

import (
	"encoding/json"
	"fmt"
)

func getAstronauts(getWebRequest GetWebRequest) (int, error) {
	url := "http://api.open-notify.org/astros.json"
	body, err := getWebRequest.FetchBytes(url)
	if err != nil {
		return 0, err
	}

	peopleResult := People{}

	if err := json.Unmarshal(body, &peopleResult); err != nil {
		return 0, err
	}

	return peopleResult.Number, err
}

func main() {
	liveClient := LiveGetWebRequest{}
	number, err := getAstronauts(liveClient)
	if err != nil {
		panic(err)
	}
	fmt.Println(number)
}
