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
	scores    map[string]string
	saveStore []string
}

func (s *StubPlayerStore) SaveStore(name string) {
	s.saveStore = append(s.saveStore, name)
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
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		store := StubPlayerStore{
			map[string]string{
				"Pepper": "20",
				"Floyd":  "10",
			},
			[]string{},
		}
		server := &PlayerServer{&store}
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		store := StubPlayerStore{
			map[string]string{
				"Pepper": "20",
				"Floyd":  "10",
			},
			[]string{},
		}
		server := &PlayerServer{&store}
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)

	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestStoreSave(t *testing.T) {
	store := StubPlayerStore{
		map[string]string{},
		[]string{},
	}
	server := &PlayerServer{&store}
	t.Run("Return accepted on post", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.saveStore) != 1 {
			t.Errorf("store different than 1")
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
