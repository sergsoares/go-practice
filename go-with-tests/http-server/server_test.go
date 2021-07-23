package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}

type StubPlayerStore struct {
	scores map[string]string
}

func (s *StubPlayerStore) GetPlayerScore(name string) string {
	score := s.scores[name]
	return score
}

func TestGetPlayers(t *testing.T) {

	t.Run("return's peppers score", func(t *testing.T) {
		store := RealPlayerStore{}
		server := &PlayerServer{&store}
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		store := StubPlayerStore{
			map[string]string{
				"Pepper": "20",
				"Floyd":  "10",
			},
		}
		server := &PlayerServer{&store}
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
