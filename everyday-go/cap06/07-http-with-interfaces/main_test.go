package main

import "testing"

type testWebRequest struct {
}

func (testWebRequest) FetchBytes(url string) ([]byte, error) {
	return []byte(`{"number":2}`), nil
}

func TestGetAstronauts(t *testing.T) {
	want := 3
	got, err := getAstronauts(testWebRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("Peopled %d, %d", want, got)
	}
}
