package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/thumbnail", thumbnailHandler)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	log.Panic(http.ListenAndServe(":8000", nil))
}

type thumbnailRequest struct {
	Url string `json:"url"`
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/photos/1", nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer response.Body.Close()

	type screenshotAPIResponse struct {
		ThumbnailUrl string `json:"thumbnailUrl"`
	}

	var apiResponse screenshotAPIResponse

	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = fmt.Fprintf(w, `{ "thumbnailUrl": "%s" }`, apiResponse.ThumbnailUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

        log.Println(apiResponse)
	fmt.Printf("Got the following url: %s\n", decoded.Url)
}
