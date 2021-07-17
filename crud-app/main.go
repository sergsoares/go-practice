package main

import (
	"crud-app/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	log.Println("Started server")

	log.Fatal(http.ListenAndServe(":8080", r))

}
