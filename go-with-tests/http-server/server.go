package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := GetPlayerScore(player)
	fmt.Fprint(w, score)
}

func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}
	if player == "Floyd" {
		return "10"
	}

	return ""
}
