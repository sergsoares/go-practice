package main

import (
	"log"
	"net/http"
	// PlayerServer "example.com/server" 
	"example.com/server" 
)

type InMemoryPlayerScore struct{}

func (i *InMemoryPlayerScore) GetPlayerScore(name string) int {
	return 999999999999
}

func main() {
	// server.HelloWorld()
	server2 := &server.PlayerServer{&InMemoryPlayerScore{}}
	log.Fatal(http.ListenAndServe(":5000", server2))
	log.Fatal("worked")
}
