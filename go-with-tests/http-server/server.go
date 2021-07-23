package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	fmt.Fprint(w, score)
}

type PlayerServer struct {
	store PlayerStore
}

type PlayerStore interface {
	GetPlayerScore(name string) string
}

type RealPlayerStore struct{}

func (r RealPlayerStore) GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}

	return ""
}
