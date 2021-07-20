package main

import "testing"

func TestFormat(t *testing.T) {
	want := "test | Aries"

	sign := sign{
		name: "test",
		sign: "Aries",
	}
	got := formatSign(sign)

	if want != got {
		t.Error("invalid format", want, got)
	}
}

func TestParseJsonToSign(t *testing.T) {
	want := sign{
		name: "Novo",
		sign: "Aries",
	}

	got := parseToSign(`{"name":"Novo","date":"2006-01-02"}`)

	if want != got {
		t.Error("Invalid parse", want, got)
	}
}

func TestMatchSignDates(t *testing.T) {
	table := []struct {
		date string
		sign string
	}{
		{"2020-02-03", "Aries"},
		{"2020-04-21", "Taurus"},
		{"2020-05-21", "Taurus"},
		{"2020-05-22", "Gemini"},
		{"2020-06-21", "Gemini"},
	}

	for _, v := range table {
		got := matchSignByDate(v.date)
		if v.sign != got {
			t.Error("Invalid match", v.date, v.sign, got)
		}

	}

}
