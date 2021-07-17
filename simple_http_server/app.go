package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Issue struct {
	Name string
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Println("HelloServer Start")
	p := Issue{Name: "Api"}

	value, err := json.Marshal(p)
	if err != nil {
		log.Fatal("Error")
		return
	}

	fmt.Fprintf(w, string(value))
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3000", nil)
}
