package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.saveScore(w)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)

}

func (p *PlayerServer) saveScore(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

type PlayerServer struct {
	store PlayerStore
}

type PlayerStore interface {
	GetPlayerScore(name string) string
	SaveStore(name string)
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
