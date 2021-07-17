package server

import (
	"fmt"
	"net/http"
	"strings"
)

func HelloWorld() {
	fmt.Print("HelloWorld")
}

type PlayerScore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerScore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.Store.GetPlayerScore(player))
}

// func GetPlayerScore(name string) string {
// 	if name == "Pepper" {
// 		return "20"
// 	}

// 	if name == "Floyd" {
// 		return "10"
// 	}

// 	return ""
// }
