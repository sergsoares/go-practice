package tidy

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthWithValidPassword(t *testing.T) {
	// Arrange
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	wantUser := "admin"
	wantPassword := "password"

	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	r.SetBasicAuth(wantUser, wantPassword)

	wantCredentials := &BasicAuthCredentials{
		User:     wantUser,
		Password: wantPassword,
	}
	decorated := DecorateWithBasicAuth(handler, wantCredentials)

	// ACT
	w := httptest.NewRecorder()
	decorated.ServeHTTP(w, r)

	// Assert
	wantCode := http.StatusOK
	if w.Code != wantCode {
		t.Fatal("Status code invalid", wantCode, w.Code)
	}

	gotAuth := w.Header().Get("WWW-Authenticate")
	wantAuth := `OK`

	// Assert
	if gotAuth != wantAuth {
		t.Fatal("auth failed", wantAuth, gotAuth)
	}

}

func TestAuthWithInvalidPasswordGives403(t *testing.T) {
	// Arrange
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	r.SetBasicAuth("admin", "123")

	wantCredentials := &BasicAuthCredentials{
		User:     "abc",
		Password: "",
	}

	handler := func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "OK") }
	decorated := DecorateWithBasicAuth(handler, wantCredentials)

	// Act
	w := httptest.NewRecorder()
	decorated.ServeHTTP(w, r)

	// Assert
	gotCode := w.Code
	wantCode := http.StatusForbidden

	if gotCode != wantCode {
		t.Fatal("Invalid error", wantCode, gotCode)
	}
}
